package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Newsroom holds the schema definition for the Newsroom entity.
type Newsroom struct {
	ent.Schema
}

// Annotations of the Newsroom.
func (Newsroom) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "newsroom"},
	}
}

// Fields of the Newsroom.
func (Newsroom) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id").Optional().Nillable(),
		field.String("header_photo_url").Optional().Nillable(),
		field.String("title").Optional().Nillable(),
		field.String("slug").Optional().Nillable().Unique(),
		field.String("category"),
		field.String("summary").Optional().Nillable(),
		field.String("body").Optional().Nillable(),
		field.String("body_rendered").Optional().Nillable(),
		field.String("link").Optional().Nillable(),
		field.Bool("force_modal"),
		field.Bool("show_timeline"),
		field.Bool("show_link"),
		field.Bool("auth_only"),
		field.Time("published_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Newsroom.
func (Newsroom) Edges() []ent.Edge {
	return []ent.Edge{}
}
