package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Hashtag holds the schema definition for the Hashtag entity.
type Hashtag struct {
	ent.Schema
}

// Annotations of the Hashtag.
func (Hashtag) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "hashtags"},
	}
}

// Fields of the Hashtag.
func (Hashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Unique(),
		field.String("slug").Unique(),
		field.Bool("can_trend").Optional().Nillable(),
		field.Bool("can_search").Optional().Nillable(),
		field.Bool("is_nsfw"),
		field.Bool("is_banned"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Uint("cached_count").Optional().Nillable(),
	}
}

// Edges of the Hashtag.
func (Hashtag) Edges() []ent.Edge {
	return []ent.Edge{}
}
