#!/usr/bin/env python3

from __future__ import annotations

import re
import subprocess
import sys
from dataclasses import dataclass
from pathlib import Path


COMPOSE_FILE_PATH = Path(".pixelfed-local/docker-compose.yml")
ENV_FILE_PATH = Path(".pixelfed-local/.env")
ENT_SCHEMA_DIR = Path("ent/schema")
TYPE_PATTERN = re.compile(r"type\s+([A-Za-z0-9_]+)\s+struct")
HIDDEN_PATTERN = re.compile(r"(?:protected|public)\s+\$hidden\s*=\s*\[(.*?)\];", re.S)
STRING_LITERAL_PATTERN = re.compile(r"""['"]([^'"]+)['"]""")
VALID_IDENTIFIER_PATTERN = re.compile(r"^[A-Za-z_][A-Za-z0-9_]*$")
RELATION_CALL_PATTERN = re.compile(
    r"return\s+\$this->"
    r"(belongsToMany|belongsTo|hasOne|hasMany|hasManyThrough|hasOneThrough|morphTo|morphMany|morphOne|morphToMany|morphedByMany)"
    r"\(\s*([A-Za-z0-9_\\]+|self)::class"
    r"(?:\s*,\s*'([^']*)')?"
    r"(?:\s*,\s*'([^']*)')?"
    r"(?:\s*,\s*'([^']*)')?"
    r"(?:\s*,\s*'([^']*)')?",
    re.S,
)
LARAVEL_MODEL_PATHS = (Path("pixelfed/app"), Path("pixelfed/app/Models"))


@dataclass
class ColumnInfo:
    name: str
    db_type: str
    nullable: bool
    key: str


@dataclass
class RelationInfo:
    name: str
    kind: str
    target: str | None
    args: tuple[str | None, str | None, str | None, str | None]
    order: int


@dataclass
class ModelInfo:
    entity_name: str
    path: Path
    relations: list[RelationInfo]


def main() -> int:
    try:
        run()
    except Exception as err:  # noqa: BLE001
        print(f"pixelfed schema sync: {err}", file=sys.stderr)
        return 1

    return 0


def run() -> None:
    env = read_env_file(ENV_FILE_PATH)
    root_password = env.get("DB_ROOT_PASSWORD", "")
    db_name = env.get("DB_DATABASE", "")
    if not root_password or not db_name:
        raise RuntimeError(f"missing DB_ROOT_PASSWORD or DB_DATABASE in {ENV_FILE_PATH}")

    tables = list_tables(root_password, db_name)
    models = load_model_infos()
    schema_files = sorted(path for path in ENT_SCHEMA_DIR.glob("*.go") if path.is_file())
    table_by_entity: dict[str, str] = {}
    columns_by_entity: dict[str, list[ColumnInfo]] = {}

    for schema_path in schema_files:
        entity_name = read_entity_name(schema_path)
        table_name = infer_table_name(entity_name, tables)
        if table_name is None:
            continue

        table_by_entity[entity_name] = table_name
        columns_by_entity[entity_name] = describe_table(root_password, db_name, table_name)

    for schema_path in schema_files:
        entity_name = read_entity_name(schema_path)
        table_name = table_by_entity.get(entity_name)
        if table_name is None:
            schema_path.unlink(missing_ok=True)
            print(f"remove {entity_name}: matching table not found", file=sys.stderr)
            continue

        columns = columns_by_entity[entity_name]
        hidden_fields = read_hidden_fields(entity_name)
        model_info = models.get(entity_name, ModelInfo(entity_name=entity_name, path=schema_path, relations=[]))
        schema_path.write_text(
            render_schema(
                entity_name,
                table_name,
                columns,
                hidden_fields,
                model_info.relations,
                models,
                columns_by_entity,
            ),
            encoding="utf-8",
        )


def read_env_file(path: Path) -> dict[str, str]:
    values: dict[str, str] = {}
    for raw_line in path.read_text(encoding="utf-8").splitlines():
        line = raw_line.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue

        key, value = line.split("=", 1)
        values[key] = value.strip().strip('"')

    return values


def list_tables(root_password: str, db_name: str) -> set[str]:
    output = run_mysql(root_password, db_name, "SHOW TABLES")
    return {line.strip() for line in output.splitlines() if line.strip()}


def read_entity_name(path: Path) -> str:
    matches = TYPE_PATTERN.search(path.read_text(encoding="utf-8"))
    if matches is None:
        raise RuntimeError(f"entity type not found in {path}")

    return matches.group(1)


def read_hidden_fields(entity_name: str) -> set[str]:
    model_path = find_laravel_model_path(entity_name)
    if model_path is None:
        return set()

    content = model_path.read_text(encoding="utf-8")
    matches = HIDDEN_PATTERN.search(content)
    if matches is None:
        return set()

    return {match.group(1) for match in STRING_LITERAL_PATTERN.finditer(matches.group(1))}


def find_laravel_model_path(entity_name: str) -> Path | None:
    for base_path in LARAVEL_MODEL_PATHS:
        candidate = base_path / f"{entity_name}.php"
        if candidate.exists():
            return candidate

    return None


def load_model_infos() -> dict[str, ModelInfo]:
    models: dict[str, ModelInfo] = {}
    for base_path in LARAVEL_MODEL_PATHS:
        for path in sorted(base_path.glob("*.php")):
            entity_name = path.stem
            models[entity_name] = ModelInfo(
                entity_name=entity_name,
                path=path,
                relations=parse_relations(entity_name, path),
            )

    return models


def parse_relations(entity_name: str, path: Path) -> list[RelationInfo]:
    relations: list[RelationInfo] = []
    for order, (method_name, body) in enumerate(iter_php_methods(path.read_text(encoding="utf-8"))):
        matches = RELATION_CALL_PATTERN.search(body)
        if matches is None:
            continue

        target = matches.group(2)
        if target == "self":
            target_name = entity_name
        else:
            target_name = target.split("\\")[-1]

        relations.append(
            RelationInfo(
                name=method_name,
                kind=matches.group(1),
                target=target_name,
                args=(matches.group(3), matches.group(4), matches.group(5), matches.group(6)),
                order=order,
            )
        )

    return relations


def iter_php_methods(content: str) -> list[tuple[str, str]]:
    methods: list[tuple[str, str]] = []
    marker = "public function "
    start = 0

    while True:
        method_start = content.find(marker, start)
        if method_start == -1:
            return methods

        name_start = method_start + len(marker)
        name_end = content.find("(", name_start)
        if name_end == -1:
            return methods

        method_name = content[name_start:name_end].strip()
        body_start = content.find("{", name_end)
        if body_start == -1:
            return methods

        depth = 0
        body_end = body_start
        while body_end < len(content):
            char = content[body_end]
            if char == "{":
                depth += 1
            elif char == "}":
                depth -= 1
                if depth == 0:
                    body_end += 1
                    break
            body_end += 1

        methods.append((method_name, content[body_start:body_end]))
        start = body_end


def infer_table_name(entity_name: str, tables: set[str]) -> str | None:
    snake = to_snake(entity_name)
    for candidate in (snake, pluralize(snake)):
        if candidate in tables:
            return candidate

    return None


def describe_table(root_password: str, db_name: str, table_name: str) -> list[ColumnInfo]:
    query = (
        "SELECT CONCAT_WS('|', COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE, IFNULL(COLUMN_KEY, '')) "
        "FROM information_schema.COLUMNS "
        f"WHERE TABLE_SCHEMA = '{db_name}' AND TABLE_NAME = '{table_name}' "
        "ORDER BY ORDINAL_POSITION"
    )
    output = run_mysql(root_password, db_name, query)

    columns: list[ColumnInfo] = []
    for line in output.splitlines():
        line = line.strip()
        if not line:
            continue

        parts = line.split("|")
        if len(parts) != 4:
            raise RuntimeError(f"unexpected column metadata for {table_name}: {line!r}")

        columns.append(
            ColumnInfo(
                name=parts[0],
                db_type=parts[1].lower(),
                nullable=parts[2].upper() == "YES",
                key=parts[3],
            )
        )

    return columns


def run_mysql(root_password: str, db_name: str, query: str) -> str:
    command = [
        "docker",
        "compose",
        "-f",
        str(COMPOSE_FILE_PATH),
        "exec",
        "-T",
        "db",
        "mysql",
        "-uroot",
        f"-p{root_password}",
        db_name,
        "-N",
        "-B",
        "-e",
        query,
    ]

    result = subprocess.run(command, capture_output=True, text=True, check=False)
    if result.returncode != 0:
        raise RuntimeError(result.stderr.strip())

    return result.stdout


def render_schema(
    entity_name: str,
    table_name: str,
    columns: list[ColumnInfo],
    hidden_fields: set[str],
    relations: list[RelationInfo],
    models: dict[str, ModelInfo],
    columns_by_entity: dict[str, list[ColumnInfo]],
) -> str:
    edges = render_edges(entity_name, relations, columns, models, columns_by_entity)
    field_type_overrides = foreign_key_type_overrides(relations, columns_by_entity)
    imports = [
        '\t"entgo.io/ent"',
        '\t"entgo.io/ent/dialect/entsql"',
        '\tentschema "entgo.io/ent/schema"',
        '\t"entgo.io/ent/schema/field"',
    ]
    if edges:
        imports.append('\t"entgo.io/ent/schema/edge"')

    lines: list[str] = [
        "package schema",
        "",
        "import (",
        *imports,
        ")",
        "",
        f"// {entity_name} holds the schema definition for the {entity_name} entity.",
        f"type {entity_name} struct {{",
        "\tent.Schema",
        "}",
        "",
        f"// Annotations of the {entity_name}.",
        f"func ({entity_name}) Annotations() []entschema.Annotation {{",
        "\treturn []entschema.Annotation{",
        f'\t\tentsql.Annotation{{Table: "{table_name}"}},',
        "\t}",
        "}",
        "",
        f"// Fields of the {entity_name}.",
        f"func ({entity_name}) Fields() []ent.Field {{",
        "\treturn []ent.Field{",
    ]

    for column in columns:
        lines.append(f"\t\t{render_field(column, hidden_fields, field_type_overrides)},")

    lines.extend(
        [
            "\t}",
            "}",
            "",
            f"// Edges of the {entity_name}.",
            f"func ({entity_name}) Edges() []ent.Edge {{",
            "\treturn []ent.Edge{",
        ]
    )

    for edge_line in edges:
        lines.append(f"\t\t{edge_line},")

    lines.extend(
        [
            "\t}",
            "}",
            "",
        ]
    )

    return "\n".join(lines)


def render_field(
    column: ColumnInfo,
    hidden_fields: set[str],
    field_type_overrides: dict[str, str],
) -> str:
    effective_db_type = field_type_overrides.get(column.name, column.db_type)
    field_name = safe_field_name(column.name)
    builder = field_builder(field_name, effective_db_type)
    if field_name != column.name:
        builder += f'.StorageKey("{column.name}")'
    if column.name in hidden_fields and supports_sensitive(effective_db_type):
        builder += ".Sensitive()"
    if column.nullable:
        builder += ".Optional().Nillable()"
    if column.key == "UNI":
        builder += ".Unique()"
    return builder


def foreign_key_type_overrides(
    relations: list[RelationInfo],
    columns_by_entity: dict[str, list[ColumnInfo]],
) -> dict[str, str]:
    overrides: dict[str, str] = {}
    for relation in relations:
        if relation.target is None:
            continue

        if relation.kind == "belongsTo":
            foreign_key = relation.args[0] or f"{to_snake(relation.name)}_id"
            owner_key = relation.args[1] or "id"
        elif relation.kind == "hasOne":
            foreign_key = relation.args[1]
            owner_key = relation.args[0] or "id"
            if foreign_key is None or foreign_key == "id":
                continue
        else:
            continue

        target_columns = columns_by_entity.get(relation.target, [])
        target_column = next((column for column in target_columns if column.name == owner_key), None)
        if target_column is None:
            continue

        overrides[foreign_key] = target_column.db_type

    return overrides


def render_edges(
    entity_name: str,
    relations: list[RelationInfo],
    columns: list[ColumnInfo],
    models: dict[str, ModelInfo],
    columns_by_entity: dict[str, list[ColumnInfo]],
) -> list[str]:
    column_by_name = {column.name: column for column in columns}
    edges: list[str] = []
    seen_signatures: set[tuple[str, str | None, str | None]] = set()

    for relation in relations:
        edge_line = render_relation_edge(
            entity_name,
            relation,
            column_by_name,
            models,
            columns_by_entity,
        )
        if edge_line is None:
            continue

        signature = edge_signature(relation)
        if signature in seen_signatures:
            continue

        seen_signatures.add(signature)
        edges.append(edge_line)

    return edges


def render_relation_edge(
    entity_name: str,
    relation: RelationInfo,
    column_by_name: dict[str, ColumnInfo],
    models: dict[str, ModelInfo],
    columns_by_entity: dict[str, list[ColumnInfo]],
) -> str | None:
    if relation.target is None:
        return None

    edge_name = resolved_edge_name(entity_name, relation, columns_by_entity)

    if relation.kind == "belongsTo":
        foreign_key = relation.args[0] or f"{to_snake(relation.name)}_id"
        foreign_key_column = column_by_name.get(foreign_key)
        if foreign_key_column is None:
            return None

        nullable = foreign_key_column.nullable
        inverse = find_inverse_relation(entity_name, relation, models)
        if inverse is not None and can_render_relation_edge(
            relation.target,
            inverse,
            columns_by_entity.get(relation.target, []),
            models,
        ):
            inverse_name = resolved_edge_name(relation.target, inverse, columns_by_entity)
            builder = (
                f'edge.From("{edge_name}", {relation.target}.Type)'
                f'.Ref("{inverse_name}").Field("{safe_field_name(foreign_key)}").Unique()'
            )
        else:
            builder = f'edge.To("{edge_name}", {relation.target}.Type).Field("{safe_field_name(foreign_key)}").Unique()'

        if not nullable:
            builder += ".Required()"
        return builder

    if relation.kind == "hasOne":
        inverse = find_inverse_relation(entity_name, relation, models)
        if inverse is not None:
            inferred_local_foreign_key = f"{to_snake(relation.name)}_id"
            if inferred_local_foreign_key in column_by_name:
                return None
            return f'edge.To("{edge_name}", {relation.target}.Type).Unique()'

        local_key = relation.args[1]
        target_key = relation.args[0] or "id"
        if local_key is None or local_key == "id" or target_key != "id":
            return None
        if local_key not in column_by_name:
            return None

        builder = f'edge.To("{edge_name}", {relation.target}.Type).Field("{safe_field_name(local_key)}").Unique()'
        if not column_by_name[local_key].nullable:
            builder += ".Required()"
        return builder

    if relation.kind == "hasMany":
        return f'edge.To("{edge_name}", {relation.target}.Type)'

    if relation.kind == "belongsToMany":
        return None

    return None


def edge_signature(relation: RelationInfo) -> tuple[str, str | None, str | None]:
    if relation.kind == "belongsTo":
        return (relation.kind, relation.target, relation.args[0] or f"{to_snake(relation.name)}_id")
    if relation.kind == "belongsToMany":
        pivot_table = relation.args[0] or ""
        owner_column = relation.args[1] or ""
        ref_column = relation.args[2] or ""
        return (relation.kind, relation.target, f"{pivot_table}:{owner_column}:{ref_column}")
    return (relation.kind, relation.target, relation.name)


def can_render_relation_edge(
    entity_name: str,
    relation: RelationInfo,
    columns: list[ColumnInfo],
    models: dict[str, ModelInfo],
) -> bool:
    column_by_name = {column.name: column for column in columns}

    if relation.kind == "belongsTo":
        foreign_key = relation.args[0] or f"{to_snake(relation.name)}_id"
        return foreign_key in column_by_name

    if relation.kind == "hasOne":
        inverse = find_inverse_relation(entity_name, relation, models)
        if inverse is not None:
            inferred_local_foreign_key = f"{to_snake(relation.name)}_id"
            return inferred_local_foreign_key not in column_by_name

        local_key = relation.args[1]
        target_key = relation.args[0] or "id"
        return local_key is not None and local_key != "id" and target_key == "id" and local_key in column_by_name

    if relation.kind == "hasMany":
        return True

    return False


def find_inverse_relation(
    entity_name: str,
    relation: RelationInfo,
    models: dict[str, ModelInfo],
) -> RelationInfo | None:
    if relation.target is None:
        return None

    target_model = models.get(relation.target)
    if target_model is None:
        return None

    candidates = [candidate for candidate in target_model.relations if candidate.target == entity_name]
    if relation.kind == "belongsTo":
        foreign_key = relation.args[0] or f"{to_snake(relation.name)}_id"
        for candidate in candidates:
            if candidate.kind not in {"hasOne", "hasMany"}:
                continue
            candidate_foreign_key = candidate.args[0] or f"{to_snake(relation.target)}_id"
            if candidate_foreign_key == foreign_key:
                return candidate
        return None

    if relation.kind in {"hasOne", "hasMany"}:
        foreign_key = relation.args[0] or f"{to_snake(entity_name)}_id"
        for candidate in candidates:
            if candidate.kind != "belongsTo":
                continue
            candidate_foreign_key = candidate.args[0] or f"{to_snake(candidate.name)}_id"
            if candidate_foreign_key == foreign_key:
                return candidate
        return None

    if relation.kind == "belongsToMany":
        pivot_table = relation.args[0]
        owner_column = relation.args[1]
        ref_column = relation.args[2]
        for candidate in candidates:
            if candidate.kind != "belongsToMany":
                continue
            if candidate.args[0] != pivot_table:
                continue
            if candidate.args[1] == ref_column and candidate.args[2] == owner_column:
                return candidate
        return None

    return None


def is_relation_owner(entity_name: str, relation: RelationInfo, inverse: RelationInfo) -> bool:
    relation_key = f"{entity_name}.{relation.name}"
    inverse_key = f"{relation.target}.{inverse.name}"
    return relation_key < inverse_key


def resolved_edge_name(
    entity_name: str,
    relation: RelationInfo,
    columns_by_entity: dict[str, list[ColumnInfo]],
) -> str:
    base_name = relation.name
    taken_names = {column.name for column in columns_by_entity.get(entity_name, [])}
    if base_name not in taken_names:
        return base_name

    candidate = f"{base_name}Relation"
    suffix = 2
    while candidate in taken_names:
        candidate = f"{base_name}Relation{suffix}"
        suffix += 1

    return candidate


def field_builder(name: str, db_type: str) -> str:
    if db_type.startswith("tinyint(1)"):
        return f'field.Bool("{name}")'
    if db_type.startswith("bigint unsigned"):
        return f'field.Uint64("{name}")'
    if db_type.startswith("bigint"):
        return f'field.Int64("{name}")'
    if (
        db_type.startswith("int unsigned")
        or db_type.startswith("mediumint unsigned")
        or db_type.startswith("smallint unsigned")
        or db_type.startswith("tinyint unsigned")
    ):
        return f'field.Uint("{name}")'
    if (
        db_type.startswith("int")
        or db_type.startswith("mediumint")
        or db_type.startswith("smallint")
        or db_type.startswith("tinyint")
    ):
        return f'field.Int("{name}")'
    if (
        db_type.startswith("decimal")
        or db_type.startswith("double")
        or db_type.startswith("float")
    ):
        return f'field.Float("{name}")'
    if (
        db_type.startswith("datetime")
        or db_type.startswith("timestamp")
        or db_type.startswith("date")
        or db_type.startswith("time")
    ):
        return f'field.Time("{name}")'
    if "blob" in db_type or db_type.startswith("binary") or db_type.startswith("varbinary"):
        return f'field.Bytes("{name}")'

    return f'field.String("{name}")'


def supports_sensitive(db_type: str) -> bool:
    if db_type.startswith("tinyint(1)"):
        return False
    if (
        db_type.startswith("bigint")
        or db_type.startswith("int")
        or db_type.startswith("mediumint")
        or db_type.startswith("smallint")
        or db_type.startswith("tinyint")
        or db_type.startswith("decimal")
        or db_type.startswith("double")
        or db_type.startswith("float")
        or db_type.startswith("datetime")
        or db_type.startswith("timestamp")
        or db_type.startswith("date")
        or db_type.startswith("time")
    ):
        return False

    return True


def safe_field_name(name: str) -> str:
    if VALID_IDENTIFIER_PATTERN.match(name):
        return name

    sanitized = name
    if sanitized and sanitized[0].isdigit():
        sanitized = normalize_leading_digits(sanitized)

    sanitized_chars: list[str] = []
    for char in sanitized:
        if char.isalnum() or char == "_":
            sanitized_chars.append(char)
        else:
            sanitized_chars.append("_")

    candidate = "".join(sanitized_chars).strip("_")
    if not candidate:
        return "field"
    if candidate[0].isdigit():
        return f"field_{candidate}"
    return candidate


def normalize_leading_digits(name: str) -> str:
    digit_words = {
        "0": "zero",
        "1": "one",
        "2": "two",
        "3": "three",
        "4": "four",
        "5": "five",
        "6": "six",
        "7": "seven",
        "8": "eight",
        "9": "nine",
    }
    parts: list[str] = []
    index = 0
    while index < len(name) and name[index].isdigit():
        parts.append(digit_words[name[index]])
        index += 1

    remainder = name[index:]
    if remainder.startswith("_"):
        return "_".join(parts) + remainder
    if remainder:
        return "_".join(parts) + "_" + remainder
    return "_".join(parts)


def to_snake(name: str) -> str:
    chars: list[str] = []
    for index, char in enumerate(name):
        if index > 0 and char.isupper():
            chars.append("_")
        chars.append(char.lower())
    return "".join(chars)


def pluralize(word: str) -> str:
    if word.endswith(("sh", "ch", "x", "s")):
        return word + "es"
    if len(word) > 1 and word.endswith("y") and word[-2] not in "aeiou":
        return word[:-1] + "ies"
    return word + "s"


if __name__ == "__main__":
    raise SystemExit(main())
