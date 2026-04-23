package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomFilterStatus holds the schema definition for the CustomFilterStatus entity.
type CustomFilterStatus struct {
	ent.Schema
}

// Annotations of the CustomFilterStatus.
func (CustomFilterStatus) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "custom_filter_statuses"},
	}
}

// Fields of the CustomFilterStatus.
func (CustomFilterStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("custom_filter_id"),
		field.Uint64("status_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CustomFilterStatus.
func (CustomFilterStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customFilter", CustomFilter.Type).Ref("statuses").Field("custom_filter_id").Unique().Required(),
		edge.To("status", Status.Type).Field("status_id").Unique().Required(),
	}
}
