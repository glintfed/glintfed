package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserEmailForgot holds the schema definition for the UserEmailForgot entity.
type UserEmailForgot struct {
	ent.Schema
}

// Annotations of the UserEmailForgot.
func (UserEmailForgot) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_email_forgots"},
	}
}

// Fields of the UserEmailForgot.
func (UserEmailForgot) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint("user_id"),
		field.String("ip_address").Optional().Nillable(),
		field.String("user_agent").Optional().Nillable(),
		field.String("referrer").Optional().Nillable(),
		field.Time("email_sent_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the UserEmailForgot.
func (UserEmailForgot) Edges() []ent.Edge {
	return []ent.Edge{}
}
