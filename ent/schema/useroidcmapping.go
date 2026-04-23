package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserOidcMapping holds the schema definition for the UserOidcMapping entity.
type UserOidcMapping struct {
	ent.Schema
}

// Annotations of the UserOidcMapping.
func (UserOidcMapping) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_oidc_mappings"},
	}
}

// Fields of the UserOidcMapping.
func (UserOidcMapping) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id"),
		field.String("oidc_id").Unique(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the UserOidcMapping.
func (UserOidcMapping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique().Required(),
	}
}
