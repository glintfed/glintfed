package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Annotations of the Profile.
func (Profile) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "profiles"},
	}
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id").Optional().Nillable(),
		field.String("domain").Optional().Nillable(),
		field.String("username").Optional().Nillable(),
		field.String("status").Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("bio").Optional().Nillable(),
		field.Bool("unlisted"),
		field.Bool("cw"),
		field.Bool("no_autolink"),
		field.String("location").Optional().Nillable(),
		field.String("website").Optional().Nillable(),
		field.String("profile_layout").Optional().Nillable(),
		field.String("header_bg").Optional().Nillable(),
		field.String("post_layout").Optional().Nillable(),
		field.Bool("is_private"),
		field.String("sharedInbox").Optional().Nillable(),
		field.String("inbox_url").Optional().Nillable(),
		field.String("outbox_url").Optional().Nillable(),
		field.String("key_id").Optional().Nillable().Unique(),
		field.String("follower_url").Optional().Nillable(),
		field.String("following_url").Optional().Nillable(),
		field.String("private_key").Sensitive().Optional().Nillable(),
		field.String("public_key").Optional().Nillable(),
		field.String("remote_url").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.Time("delete_after").Optional().Nillable(),
		field.Bool("is_suggestable"),
		field.Time("last_fetched_at").Optional().Nillable(),
		field.Uint("status_count").Optional().Nillable(),
		field.Uint("followers_count").Optional().Nillable(),
		field.Uint("following_count").Optional().Nillable(),
		field.String("webfinger").Optional().Nillable().Unique(),
		field.String("avatar_url").Optional().Nillable(),
		field.Time("last_status_at").Optional().Nillable(),
		field.Uint64("moved_to_profile_id").Optional().Nillable(),
		field.Bool("indexable"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique(),
		edge.To("statuses", Status.Type),
		edge.To("likes", Like.Type),
		edge.To("avatar", Avatar.Type).Unique(),
		edge.To("reports", Report.Type),
		edge.To("media", Media.Type),
		edge.To("circles", Circle.Type),
		edge.To("hashtagFollowing", HashtagFollow.Type),
		edge.To("collections", Collection.Type),
		edge.To("stories", Story.Type),
		edge.To("reported", Report.Type),
		edge.To("aliases", ProfileAlias.Type),
	}
}
