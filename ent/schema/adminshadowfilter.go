package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AdminShadowFilter holds the schema definition for the AdminShadowFilter entity.
type AdminShadowFilter struct {
	ent.Schema
}

// Annotations of the AdminShadowFilter.
func (AdminShadowFilter) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "admin_shadow_filters"},
	}
}

// Fields of the AdminShadowFilter.
func (AdminShadowFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("admin_id").Optional().Nillable(),
		field.String("item_type"),
		field.Uint64("item_id"),
		field.Bool("is_local"),
		field.String("note").Optional().Nillable(),
		field.Bool("active"),
		field.String("history").Optional().Nillable(),
		field.String("ruleset").Optional().Nillable(),
		field.Bool("prevent_ap_fanout"),
		field.Bool("prevent_new_dms"),
		field.Bool("ignore_reports"),
		field.Bool("ignore_mentions"),
		field.Bool("ignore_links"),
		field.Bool("ignore_hashtags"),
		field.Bool("hide_from_public_feeds"),
		field.Bool("hide_from_tag_feeds"),
		field.Bool("hide_embeds"),
		field.Bool("hide_from_story_carousel"),
		field.Bool("hide_from_search_autocomplete"),
		field.Bool("hide_from_search"),
		field.Bool("requires_login"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the AdminShadowFilter.
func (AdminShadowFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("item_id").Unique().Required(),
	}
}
