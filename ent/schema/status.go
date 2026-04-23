package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Annotations of the Status.
func (Status) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "statuses"},
	}
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("uri").Optional().Nillable().Unique(),
		field.String("caption"),
		field.String("rendered"),
		field.Uint64("profile_id").Optional().Nillable(),
		field.String("type").Optional().Nillable(),
		field.Uint64("in_reply_to_id").Optional().Nillable(),
		field.Uint64("reblog_of_id").Optional().Nillable(),
		field.String("url").Optional().Nillable(),
		field.Bool("is_nsfw"),
		field.String("scope"),
		field.String("visibility"),
		field.Bool("reply"),
		field.Uint64("likes_count"),
		field.Uint64("reblogs_count"),
		field.String("language").Optional().Nillable(),
		field.Uint64("conversation_id").Optional().Nillable(),
		field.Bool("local"),
		field.Uint64("application_id").Optional().Nillable(),
		field.Uint64("in_reply_to_profile_id").Optional().Nillable(),
		field.String("entities").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.String("cw_summary").Optional().Nillable(),
		field.Uint("reply_count").Optional().Nillable(),
		field.Bool("comments_disabled"),
		field.Uint64("place_id").Optional().Nillable(),
		field.String("object_url").Optional().Nillable().Unique(),
		field.Time("edited_at").Optional().Nillable(),
		field.Bool("trendable").Optional().Nillable(),
		field.String("media_ids").Optional().Nillable(),
		field.Int("pinned_order").Optional().Nillable(),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("statuses").Field("profile_id").Unique(),
		edge.To("media", Media.Type),
		edge.To("firstMedia", Media.Type),
		edge.To("likes", Like.Type),
		edge.To("comments", Status.Type),
		edge.To("shares", Status.Type),
		edge.From("place", Place.Type).Ref("posts").Field("place_id").Unique(),
		edge.To("directMessage", DirectMessage.Type).Unique(),
		edge.To("edits", StatusEdit.Type),
	}
}
