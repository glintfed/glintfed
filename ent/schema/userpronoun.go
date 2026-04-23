package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserPronoun holds the schema definition for the UserPronoun entity.
type UserPronoun struct {
	ent.Schema
}

// Annotations of the UserPronoun.
func (UserPronoun) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_pronouns"},
	}
}

// Fields of the UserPronoun.
func (UserPronoun) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint("user_id").Optional().Nillable().Unique(),
		field.Int64("profile_id").Unique(),
		field.String("pronouns").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the UserPronoun.
func (UserPronoun) Edges() []ent.Edge {
	return []ent.Edge{}
}
