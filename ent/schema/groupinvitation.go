package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupInvitation holds the schema definition for the GroupInvitation entity.
type GroupInvitation struct {
	ent.Schema
}

// Annotations of the GroupInvitation.
func (GroupInvitation) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_invitations"},
	}
}

// Fields of the GroupInvitation.
func (GroupInvitation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("from_profile_id"),
		field.Uint64("to_profile_id"),
		field.String("role").Optional().Nillable(),
		field.Bool("to_local"),
		field.Bool("from_local"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupInvitation.
func (GroupInvitation) Edges() []ent.Edge {
	return []ent.Edge{}
}
