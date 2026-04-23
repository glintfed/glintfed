package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserInvite holds the schema definition for the UserInvite entity.
type UserInvite struct {
	ent.Schema
}

// Annotations of the UserInvite.
func (UserInvite) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_invites"},
	}
}

// Fields of the UserInvite.
func (UserInvite) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id"),
		field.Uint64("profile_id"),
		field.String("email").Unique(),
		field.String("message").Optional().Nillable(),
		field.String("key"),
		field.String("token"),
		field.Time("valid_until").Optional().Nillable(),
		field.Time("used_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the UserInvite.
func (UserInvite) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sender", Profile.Type).Field("profile_id").Unique().Required(),
	}
}
