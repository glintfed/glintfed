package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PulseAggregate holds the schema definition for the PulseAggregate entity.
type PulseAggregate struct {
	ent.Schema
}

// Annotations of the PulseAggregate.
func (PulseAggregate) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "pulse_aggregates"},
	}
}

// Fields of the PulseAggregate.
func (PulseAggregate) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint("bucket"),
		field.Uint("period"),
		field.String("type"),
		field.String("key"),
		field.Bytes("key_hash").Optional().Nillable(),
		field.String("aggregate"),
		field.Float("value"),
		field.Uint("count").Optional().Nillable(),
	}
}

// Edges of the PulseAggregate.
func (PulseAggregate) Edges() []ent.Edge {
	return []ent.Edge{}
}
