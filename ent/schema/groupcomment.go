package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupComment holds the schema definition for the GroupComment entity.
type GroupComment struct {
	ent.Schema
}

// Annotations of the GroupComment.
func (GroupComment) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_comments"},
	}
}

// Fields of the GroupComment.
func (GroupComment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id").Optional().Nillable(),
		field.Uint64("status_id").Optional().Nillable(),
		field.Uint64("in_reply_to_id").Optional().Nillable(),
		field.String("remote_url").Optional().Nillable().Unique(),
		field.String("caption").Optional().Nillable(),
		field.Bool("is_nsfw"),
		field.String("visibility").Optional().Nillable(),
		field.Uint("likes_count"),
		field.Uint("replies_count"),
		field.String("cw_summary").Optional().Nillable(),
		field.String("media_ids").Optional().Nillable(),
		field.String("status").Optional().Nillable(),
		field.String("type").Optional().Nillable(),
		field.Bool("local"),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupComment.
func (GroupComment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique(),
	}
}
