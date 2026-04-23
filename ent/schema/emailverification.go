package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EmailVerification holds the schema definition for the EmailVerification entity.
type EmailVerification struct {
	ent.Schema
}

// Annotations of the EmailVerification.
func (EmailVerification) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "email_verifications"},
	}
}

// Fields of the EmailVerification.
func (EmailVerification) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id"),
		field.String("email").Optional().Nillable(),
		field.String("user_token"),
		field.String("random_token"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the EmailVerification.
func (EmailVerification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique().Required(),
	}
}
