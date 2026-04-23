package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupActivityGraph holds the schema definition for the GroupActivityGraph entity.
type GroupActivityGraph struct {
	ent.Schema
}

// Annotations of the GroupActivityGraph.
func (GroupActivityGraph) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_activity_graphs"},
	}
}

// Fields of the GroupActivityGraph.
func (GroupActivityGraph) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Int64("instance_id").Optional().Nillable(),
		field.Int64("actor_id").Optional().Nillable(),
		field.String("verb").Optional().Nillable(),
		field.String("id_url").Optional().Nillable().Unique(),
		field.String("payload").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupActivityGraph.
func (GroupActivityGraph) Edges() []ent.Edge {
	return []ent.Edge{}
}
