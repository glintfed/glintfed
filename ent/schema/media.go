package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Annotations of the Media.
func (Media) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "media"},
	}
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.Uint64("status_id").Optional().Nillable(),
		field.Uint64("profile_id").Optional().Nillable(),
		field.Uint64("user_id").Optional().Nillable(),
		field.Bool("is_nsfw"),
		field.Bool("remote_media"),
		field.String("original_sha256").Optional().Nillable(),
		field.String("optimized_sha256").Optional().Nillable(),
		field.String("media_path"),
		field.String("thumbnail_path").Optional().Nillable(),
		field.String("cdn_url").Optional().Nillable(),
		field.String("optimized_url").Optional().Nillable(),
		field.String("thumbnail_url").Optional().Nillable(),
		field.String("remote_url").Optional().Nillable(),
		field.String("caption").Optional().Nillable(),
		field.String("hls_path").Optional().Nillable(),
		field.Uint("order"),
		field.String("mime").Optional().Nillable(),
		field.Uint("size").Optional().Nillable(),
		field.String("orientation").Optional().Nillable(),
		field.String("filter_name").Optional().Nillable(),
		field.String("filter_class").Optional().Nillable(),
		field.String("license").Optional().Nillable(),
		field.Time("processed_at").Optional().Nillable(),
		field.Time("hls_transcoded_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.String("key").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Int("version"),
		field.String("blurhash").Optional().Nillable(),
		field.String("srcset").Optional().Nillable(),
		field.Uint("width").Optional().Nillable(),
		field.Uint("height").Optional().Nillable(),
		field.Bool("skip_optimize").Optional().Nillable(),
		field.Time("replicated_at").Optional().Nillable(),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("status", Status.Type).Ref("media").Field("status_id").Unique(),
		edge.From("profile", Profile.Type).Ref("media").Field("profile_id").Unique(),
	}
}
