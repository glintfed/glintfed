package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// DiscoverCategoryHashtag holds the schema definition for the DiscoverCategoryHashtag entity.
type DiscoverCategoryHashtag struct {
	ent.Schema
}

// Annotations of the DiscoverCategoryHashtag.
func (DiscoverCategoryHashtag) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "discover_category_hashtags"},
	}
}

// Fields of the DiscoverCategoryHashtag.
func (DiscoverCategoryHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("discover_category_id"),
		field.Uint64("hashtag_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the DiscoverCategoryHashtag.
func (DiscoverCategoryHashtag) Edges() []ent.Edge {
	return []ent.Edge{}
}
