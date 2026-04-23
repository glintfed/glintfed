package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthRefreshToken holds the schema definition for the OauthRefreshToken entity.
type OauthRefreshToken struct {
	ent.Schema
}

// Annotations of the OauthRefreshToken.
func (OauthRefreshToken) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "oauth_refresh_tokens"},
	}
}

// Fields of the OauthRefreshToken.
func (OauthRefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("access_token_id"),
		field.Bool("revoked"),
		field.Time("expires_at").Optional().Nillable(),
	}
}

// Edges of the OauthRefreshToken.
func (OauthRefreshToken) Edges() []ent.Edge {
	return []ent.Edge{}
}
