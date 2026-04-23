package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FollowRequest holds the schema definition for the FollowRequest entity.
type FollowRequest struct {
	ent.Schema
}

// Annotations of the FollowRequest.
func (FollowRequest) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "follow_requests"},
	}
}

// Fields of the FollowRequest.
func (FollowRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("follower_id"),
		field.Uint64("following_id"),
		field.String("activity").Optional().Nillable(),
		field.Bool("is_rejected"),
		field.Bool("is_local"),
		field.Time("handled_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the FollowRequest.
func (FollowRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actor", Profile.Type).Field("follower_id").Unique().Required(),
		edge.To("following", Profile.Type).Field("following_id").Unique().Required(),
	}
}
