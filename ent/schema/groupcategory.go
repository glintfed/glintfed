package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupCategory holds the schema definition for the GroupCategory entity.
type GroupCategory struct {
	ent.Schema
}

// Annotations of the GroupCategory.
func (GroupCategory) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_categories"},
	}
}

// Fields of the GroupCategory.
func (GroupCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Unique(),
		field.String("slug").Unique(),
		field.Bool("active"),
		field.Uint("order").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupCategory.
func (GroupCategory) Edges() []ent.Edge {
	return []ent.Edge{}
}
