package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ImportData holds the schema definition for the ImportData entity.
type ImportData struct {
	ent.Schema
}

// Annotations of the ImportData.
func (ImportData) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "import_datas"},
	}
}

// Fields of the ImportData.
func (ImportData) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.Uint64("profile_id"),
		field.Uint64("job_id").Optional().Nillable(),
		field.String("service"),
		field.String("path").Optional().Nillable(),
		field.Uint("stage"),
		field.String("original_name").Optional().Nillable(),
		field.Bool("import_accepted").Optional().Nillable(),
		field.Time("completed_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ImportData.
func (ImportData) Edges() []ent.Edge {
	return []ent.Edge{}
}
