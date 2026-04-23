package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Annotations of the Notification.
func (Notification) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "notifications"},
	}
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.Uint64("actor_id").Optional().Nillable(),
		field.Uint64("item_id").Optional().Nillable(),
		field.String("item_type").Optional().Nillable(),
		field.String("action").Optional().Nillable(),
		field.Time("read_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actor", Profile.Type).Field("actor_id").Unique(),
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
		edge.To("status", Status.Type).Field("item_id").Unique(),
	}
}
