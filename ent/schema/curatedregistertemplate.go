package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CuratedRegisterTemplate holds the schema definition for the CuratedRegisterTemplate entity.
type CuratedRegisterTemplate struct {
	ent.Schema
}

// Annotations of the CuratedRegisterTemplate.
func (CuratedRegisterTemplate) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "curated_register_templates"},
	}
}

// Fields of the CuratedRegisterTemplate.
func (CuratedRegisterTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.String("content").Optional().Nillable(),
		field.Bool("is_active"),
		field.Uint("order"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CuratedRegisterTemplate.
func (CuratedRegisterTemplate) Edges() []ent.Edge {
	return []ent.Edge{}
}
