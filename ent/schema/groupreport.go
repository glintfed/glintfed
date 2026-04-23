package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupReport holds the schema definition for the GroupReport entity.
type GroupReport struct {
	ent.Schema
}

// Annotations of the GroupReport.
func (GroupReport) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_reports"},
	}
}

// Fields of the GroupReport.
func (GroupReport) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.String("type").Optional().Nillable(),
		field.String("item_type").Optional().Nillable(),
		field.String("item_id").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Bool("open"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupReport.
func (GroupReport) Edges() []ent.Edge {
	return []ent.Edge{}
}
