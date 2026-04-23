package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Follower holds the schema definition for the Follower entity.
type Follower struct {
	ent.Schema
}

// Annotations of the Follower.
func (Follower) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "followers"},
	}
}

// Fields of the Follower.
func (Follower) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.Uint64("profile_id"),
		field.Uint64("following_id"),
		field.Bool("local_profile"),
		field.Bool("local_following"),
		field.Bool("show_reblogs"),
		field.Bool("notify"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Follower.
func (Follower) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actor", Profile.Type).Field("profile_id").Unique().Required(),
		edge.To("target", Profile.Type).Field("following_id").Unique().Required(),
	}
}
