package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupMedia holds the schema definition for the GroupMedia entity.
type GroupMedia struct {
	ent.Schema
}

// Annotations of the GroupMedia.
func (GroupMedia) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_media"},
	}
}

// Fields of the GroupMedia.
func (GroupMedia) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.Uint64("status_id").Optional().Nillable(),
		field.String("media_path").Unique(),
		field.String("thumbnail_url").Optional().Nillable(),
		field.String("cdn_url").Optional().Nillable(),
		field.String("url").Optional().Nillable(),
		field.String("mime").Optional().Nillable(),
		field.Uint("size").Optional().Nillable(),
		field.String("cw_summary").Optional().Nillable(),
		field.String("license").Optional().Nillable(),
		field.String("blurhash").Optional().Nillable(),
		field.Uint("order"),
		field.Uint("width").Optional().Nillable(),
		field.Uint("height").Optional().Nillable(),
		field.Bool("local_user"),
		field.Bool("is_cached"),
		field.Bool("is_comment"),
		field.String("metadata").Optional().Nillable(),
		field.String("version"),
		field.Bool("skip_optimize"),
		field.Time("processed_at").Optional().Nillable(),
		field.Time("thumbnail_generated").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupMedia.
func (GroupMedia) Edges() []ent.Edge {
	return []ent.Edge{}
}
