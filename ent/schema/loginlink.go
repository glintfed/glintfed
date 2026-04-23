package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// LoginLink holds the schema definition for the LoginLink entity.
type LoginLink struct {
	ent.Schema
}

// Annotations of the LoginLink.
func (LoginLink) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "login_links"},
	}
}

// Fields of the LoginLink.
func (LoginLink) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("key"),
		field.String("secret"),
		field.Uint("user_id"),
		field.String("ip").Optional().Nillable(),
		field.String("user_agent").Optional().Nillable(),
		field.String("meta").Optional().Nillable(),
		field.Time("revoked_at").Optional().Nillable(),
		field.Time("resent_at").Optional().Nillable(),
		field.Time("used_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the LoginLink.
func (LoginLink) Edges() []ent.Edge {
	return []ent.Edge{}
}
