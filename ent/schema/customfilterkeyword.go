package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomFilterKeyword holds the schema definition for the CustomFilterKeyword entity.
type CustomFilterKeyword struct {
	ent.Schema
}

// Annotations of the CustomFilterKeyword.
func (CustomFilterKeyword) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "custom_filter_keywords"},
	}
}

// Fields of the CustomFilterKeyword.
func (CustomFilterKeyword) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("custom_filter_id"),
		field.String("keyword"),
		field.Bool("whole_word"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CustomFilterKeyword.
func (CustomFilterKeyword) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customFilter", CustomFilter.Type).Ref("keywords").Field("custom_filter_id").Unique().Required(),
	}
}
