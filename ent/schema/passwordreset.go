package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PasswordReset holds the schema definition for the PasswordReset entity.
type PasswordReset struct {
	ent.Schema
}

// Annotations of the PasswordReset.
func (PasswordReset) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "password_resets"},
	}
}

// Fields of the PasswordReset.
func (PasswordReset) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("token"),
		field.Time("created_at").Optional().Nillable(),
	}
}

// Edges of the PasswordReset.
func (PasswordReset) Edges() []ent.Edge {
	return []ent.Edge{}
}
