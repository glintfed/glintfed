package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DiscoverCategory holds the schema definition for the DiscoverCategory entity.
type DiscoverCategory struct {
	ent.Schema
}

// Annotations of the DiscoverCategory.
func (DiscoverCategory) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "discover_categories"},
	}
}

// Fields of the DiscoverCategory.
func (DiscoverCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Optional().Nillable(),
		field.String("slug").Unique(),
		field.Bool("active"),
		field.Uint("order"),
		field.Uint("media_id").Optional().Nillable().Unique(),
		field.Bool("no_nsfw"),
		field.Bool("local_only"),
		field.Bool("public_only"),
		field.Bool("photos_only"),
		field.Time("active_until").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the DiscoverCategory.
func (DiscoverCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("media", Media.Type).Field("media_id").Unique(),
		edge.To("items", DiscoverCategoryHashtag.Type),
	}
}
