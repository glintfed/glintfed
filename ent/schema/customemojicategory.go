package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CustomEmojiCategory holds the schema definition for the CustomEmojiCategory entity.
type CustomEmojiCategory struct {
	ent.Schema
}

// Annotations of the CustomEmojiCategory.
func (CustomEmojiCategory) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "custom_emoji_categories"},
	}
}

// Fields of the CustomEmojiCategory.
func (CustomEmojiCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Unique(),
		field.Bool("disabled"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CustomEmojiCategory.
func (CustomEmojiCategory) Edges() []ent.Edge {
	return []ent.Edge{}
}
