package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Annotations of the Page.
func (Page) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "pages"},
	}
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("root").Optional().Nillable(),
		field.String("slug").Optional().Nillable().Unique(),
		field.String("title").Optional().Nillable(),
		field.Uint("category_id").Optional().Nillable(),
		field.String("content").Optional().Nillable(),
		field.String("template"),
		field.Bool("active"),
		field.Bool("cached"),
		field.Time("active_until").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{}
}
