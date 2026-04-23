package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StatusArchived holds the schema definition for the StatusArchived entity.
type StatusArchived struct {
	ent.Schema
}

// Annotations of the StatusArchived.
func (StatusArchived) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "status_archiveds"},
	}
}

// Fields of the StatusArchived.
func (StatusArchived) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("status_id"),
		field.Uint64("profile_id"),
		field.String("original_scope").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the StatusArchived.
func (StatusArchived) Edges() []ent.Edge {
	return []ent.Edge{}
}
