package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthAccessToken holds the schema definition for the OauthAccessToken entity.
type OauthAccessToken struct {
	ent.Schema
}

// Annotations of the OauthAccessToken.
func (OauthAccessToken) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "oauth_access_tokens"},
	}
}

// Fields of the OauthAccessToken.
func (OauthAccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Uint64("user_id").Optional().Nillable(),
		field.Uint64("client_id"),
		field.String("name").Optional().Nillable(),
		field.String("scopes").Optional().Nillable(),
		field.Bool("revoked"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("expires_at").Optional().Nillable(),
	}
}

// Edges of the OauthAccessToken.
func (OauthAccessToken) Edges() []ent.Edge {
	return []ent.Edge{}
}
