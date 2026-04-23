package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthPersonalAccessClient holds the schema definition for the OauthPersonalAccessClient entity.
type OauthPersonalAccessClient struct {
	ent.Schema
}

// Annotations of the OauthPersonalAccessClient.
func (OauthPersonalAccessClient) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "oauth_personal_access_clients"},
	}
}

// Fields of the OauthPersonalAccessClient.
func (OauthPersonalAccessClient) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("client_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the OauthPersonalAccessClient.
func (OauthPersonalAccessClient) Edges() []ent.Edge {
	return []ent.Edge{}
}
