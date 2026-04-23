package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OauthClient holds the schema definition for the OauthClient entity.
type OauthClient struct {
	ent.Schema
}

// Annotations of the OauthClient.
func (OauthClient) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "oauth_clients"},
	}
}

// Fields of the OauthClient.
func (OauthClient) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id").Optional().Nillable(),
		field.String("name"),
		field.String("secret").Optional().Nillable(),
		field.String("provider").Optional().Nillable(),
		field.String("redirect"),
		field.Bool("personal_access_client"),
		field.Bool("password_client"),
		field.Bool("revoked"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the OauthClient.
func (OauthClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique(),
	}
}
