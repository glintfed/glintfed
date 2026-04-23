package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupLimit holds the schema definition for the GroupLimit entity.
type GroupLimit struct {
	ent.Schema
}

// Annotations of the GroupLimit.
func (GroupLimit) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_limits"},
	}
}

// Fields of the GroupLimit.
func (GroupLimit) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.String("limits").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupLimit.
func (GroupLimit) Edges() []ent.Edge {
	return []ent.Edge{}
}
