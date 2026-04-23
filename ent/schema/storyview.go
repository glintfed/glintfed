package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StoryView holds the schema definition for the StoryView entity.
type StoryView struct {
	ent.Schema
}

// Annotations of the StoryView.
func (StoryView) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "story_views"},
	}
}

// Fields of the StoryView.
func (StoryView) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("story_id"),
		field.Uint64("profile_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the StoryView.
func (StoryView) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("story", Story.Type).Ref("views").Field("story_id").Unique().Required(),
	}
}
