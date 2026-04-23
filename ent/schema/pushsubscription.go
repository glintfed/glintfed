package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PushSubscription holds the schema definition for the PushSubscription entity.
type PushSubscription struct {
	ent.Schema
}

// Annotations of the PushSubscription.
func (PushSubscription) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "push_subscriptions"},
	}
}

// Fields of the PushSubscription.
func (PushSubscription) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("subscribable_type"),
		field.Uint64("subscribable_id"),
		field.String("endpoint").Unique(),
		field.String("public_key").Optional().Nillable(),
		field.String("auth_token").Optional().Nillable(),
		field.String("content_encoding").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the PushSubscription.
func (PushSubscription) Edges() []ent.Edge {
	return []ent.Edge{}
}
