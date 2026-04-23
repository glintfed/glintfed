package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Migration holds the schema definition for the Migration entity.
type Migration struct {
	ent.Schema
}

// Annotations of the Migration.
func (Migration) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "migrations"},
	}
}

// Fields of the Migration.
func (Migration) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("migration"),
		field.Int("batch"),
	}
}

// Edges of the Migration.
func (Migration) Edges() []ent.Edge {
	return []ent.Edge{}
}
