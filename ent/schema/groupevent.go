package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupEvent holds the schema definition for the GroupEvent entity.
type GroupEvent struct {
	ent.Schema
}

// Annotations of the GroupEvent.
func (GroupEvent) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_events"},
	}
}

// Fields of the GroupEvent.
func (GroupEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id").Optional().Nillable(),
		field.Uint64("profile_id").Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("type"),
		field.String("tags").Optional().Nillable(),
		field.String("location").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Bool("open"),
		field.Bool("comments_open"),
		field.Bool("show_guest_list"),
		field.Time("start_at").Optional().Nillable(),
		field.Time("end_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupEvent.
func (GroupEvent) Edges() []ent.Edge {
	return []ent.Edge{}
}
