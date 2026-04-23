package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RemoteAuth holds the schema definition for the RemoteAuth entity.
type RemoteAuth struct {
	ent.Schema
}

// Annotations of the RemoteAuth.
func (RemoteAuth) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "remote_auths"},
	}
}

// Fields of the RemoteAuth.
func (RemoteAuth) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("software").Optional().Nillable(),
		field.String("domain").Optional().Nillable(),
		field.String("webfinger").Optional().Nillable().Unique(),
		field.Uint("instance_id").Optional().Nillable(),
		field.Uint("user_id").Optional().Nillable().Unique(),
		field.Uint("client_id").Optional().Nillable(),
		field.String("ip_address").Optional().Nillable(),
		field.String("bearer_token").Optional().Nillable(),
		field.String("verify_credentials").Optional().Nillable(),
		field.Time("last_successful_login_at").Optional().Nillable(),
		field.Time("last_verify_credentials_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the RemoteAuth.
func (RemoteAuth) Edges() []ent.Edge {
	return []ent.Edge{}
}
