package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupPost holds the schema definition for the GroupPost entity.
type GroupPost struct {
	ent.Schema
}

// Annotations of the GroupPost.
func (GroupPost) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_posts"},
	}
}

// Fields of the GroupPost.
func (GroupPost) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id").Optional().Nillable(),
		field.String("type").Optional().Nillable(),
		field.String("remote_url").Optional().Nillable().Unique(),
		field.Uint("reply_count").Optional().Nillable(),
		field.String("status").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.String("caption").Optional().Nillable(),
		field.String("visibility").Optional().Nillable(),
		field.Bool("is_nsfw"),
		field.Uint("likes_count"),
		field.String("cw_summary").Optional().Nillable(),
		field.String("media_ids").Optional().Nillable(),
		field.Bool("comments_disabled"),
	}
}

// Edges of the GroupPost.
func (GroupPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type).Field("group_id").Unique().Required(),
		edge.To("profile", Profile.Type).Field("profile_id").Unique(),
	}
}
