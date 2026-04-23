package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupInteraction holds the schema definition for the GroupInteraction entity.
type GroupInteraction struct {
	ent.Schema
}

// Annotations of the GroupInteraction.
func (GroupInteraction) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_interactions"},
	}
}

// Fields of the GroupInteraction.
func (GroupInteraction) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.String("type").Optional().Nillable(),
		field.String("item_type").Optional().Nillable(),
		field.String("item_id").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupInteraction.
func (GroupInteraction) Edges() []ent.Edge {
	return []ent.Edge{}
}
