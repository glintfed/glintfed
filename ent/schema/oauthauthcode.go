package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthAuthCode holds the schema definition for the OauthAuthCode entity.
type OauthAuthCode struct {
	ent.Schema
}

// Annotations of the OauthAuthCode.
func (OauthAuthCode) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "oauth_auth_codes"},
	}
}

// Fields of the OauthAuthCode.
func (OauthAuthCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Uint64("user_id"),
		field.Uint64("client_id"),
		field.String("scopes").Optional().Nillable(),
		field.Bool("revoked"),
		field.Time("expires_at").Optional().Nillable(),
	}
}

// Edges of the OauthAuthCode.
func (OauthAuthCode) Edges() []ent.Edge {
	return []ent.Edge{}
}
