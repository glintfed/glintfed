package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupRole holds the schema definition for the GroupRole entity.
type GroupRole struct {
	ent.Schema
}

// Annotations of the GroupRole.
func (GroupRole) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_roles"},
	}
}

// Fields of the GroupRole.
func (GroupRole) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.String("name"),
		field.String("slug").Optional().Nillable(),
		field.String("abilities").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupRole.
func (GroupRole) Edges() []ent.Edge {
	return []ent.Edge{}
}
