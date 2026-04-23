package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StatusView holds the schema definition for the StatusView entity.
type StatusView struct {
	ent.Schema
}

// Annotations of the StatusView.
func (StatusView) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "status_views"},
	}
}

// Fields of the StatusView.
func (StatusView) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("status_id").Optional().Nillable(),
		field.Uint64("status_profile_id").Optional().Nillable(),
		field.Uint64("profile_id").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the StatusView.
func (StatusView) Edges() []ent.Edge {
	return []ent.Edge{}
}
