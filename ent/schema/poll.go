package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Annotations of the Poll.
func (Poll) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "polls"},
	}
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("story_id").Optional().Nillable(),
		field.Uint64("status_id").Optional().Nillable(),
		field.Uint64("group_id").Optional().Nillable(),
		field.Uint64("profile_id"),
		field.String("poll_options").Optional().Nillable(),
		field.String("cached_tallies").Optional().Nillable(),
		field.Bool("multiple"),
		field.Bool("hide_totals"),
		field.Uint("votes_count"),
		field.Time("last_fetched_at").Optional().Nillable(),
		field.Time("expires_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Poll.
func (Poll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("votes", PollVote.Type),
	}
}
