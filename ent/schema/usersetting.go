package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserSetting holds the schema definition for the UserSetting entity.
type UserSetting struct {
	ent.Schema
}

// Annotations of the UserSetting.
func (UserSetting) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_settings"},
	}
}

// Fields of the UserSetting.
func (UserSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id").Unique(),
		field.String("role"),
		field.Bool("crawlable"),
		field.Bool("show_guests"),
		field.Bool("show_discover"),
		field.Bool("public_dm"),
		field.Bool("hide_cw_search"),
		field.Bool("hide_blocked_search"),
		field.Bool("always_show_cw"),
		field.Bool("compose_media_descriptions"),
		field.Bool("reduce_motion"),
		field.Bool("optimize_screen_reader"),
		field.Bool("high_contrast_mode"),
		field.Bool("video_autoplay"),
		field.Bool("send_email_new_follower"),
		field.Bool("send_email_new_follower_request"),
		field.Bool("send_email_on_share"),
		field.Bool("send_email_on_like"),
		field.Bool("send_email_on_mention"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Bool("show_profile_followers"),
		field.Bool("show_profile_follower_count"),
		field.Bool("show_profile_following"),
		field.Bool("show_profile_following_count"),
		field.String("compose_settings").Optional().Nillable(),
		field.String("other").Optional().Nillable(),
		field.Bool("show_atom"),
	}
}

// Edges of the UserSetting.
func (UserSetting) Edges() []ent.Edge {
	return []ent.Edge{}
}
