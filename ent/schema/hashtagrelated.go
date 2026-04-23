package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// HashtagRelated holds the schema definition for the HashtagRelated entity.
type HashtagRelated struct {
	ent.Schema
}

// Annotations of the HashtagRelated.
func (HashtagRelated) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "hashtag_related"},
	}
}

// Fields of the HashtagRelated.
func (HashtagRelated) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("hashtag_id").Unique(),
		field.String("related_tags").Optional().Nillable(),
		field.Uint64("agg_score").Optional().Nillable(),
		field.Time("last_calculated_at").Optional().Nillable(),
		field.Time("last_moderated_at").Optional().Nillable(),
		field.Bool("skip_refresh"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the HashtagRelated.
func (HashtagRelated) Edges() []ent.Edge {
	return []ent.Edge{}
}
