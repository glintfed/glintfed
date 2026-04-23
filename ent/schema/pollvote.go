package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PollVote holds the schema definition for the PollVote entity.
type PollVote struct {
	ent.Schema
}

// Annotations of the PollVote.
func (PollVote) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "poll_votes"},
	}
}

// Fields of the PollVote.
func (PollVote) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("story_id").Optional().Nillable(),
		field.Uint64("status_id").Optional().Nillable(),
		field.Uint64("profile_id"),
		field.Uint64("poll_id"),
		field.Uint("choice"),
		field.String("uri").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the PollVote.
func (PollVote) Edges() []ent.Edge {
	return []ent.Edge{}
}
