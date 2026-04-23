package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CustomEmoji holds the schema definition for the CustomEmoji entity.
type CustomEmoji struct {
	ent.Schema
}

// Annotations of the CustomEmoji.
func (CustomEmoji) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "custom_emoji"},
	}
}

// Fields of the CustomEmoji.
func (CustomEmoji) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("shortcode"),
		field.String("media_path").Optional().Nillable(),
		field.String("domain").Optional().Nillable(),
		field.Bool("disabled"),
		field.String("uri").Optional().Nillable(),
		field.String("image_remote_url").Optional().Nillable(),
		field.Uint("category_id").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CustomEmoji.
func (CustomEmoji) Edges() []ent.Edge {
	return []ent.Edge{}
}
