package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserRoles holds the schema definition for the UserRoles entity.
type UserRoles struct {
	ent.Schema
}

// Annotations of the UserRoles.
func (UserRoles) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_roles"},
	}
}

// Fields of the UserRoles.
func (UserRoles) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id").Optional().Nillable().Unique(),
		field.Uint64("user_id").Unique(),
		field.String("roles").Optional().Nillable(),
		field.String("meta").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the UserRoles.
func (UserRoles) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique().Required(),
	}
}
