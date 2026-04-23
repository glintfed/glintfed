package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PulseValue holds the schema definition for the PulseValue entity.
type PulseValue struct {
	ent.Schema
}

// Annotations of the PulseValue.
func (PulseValue) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "pulse_values"},
	}
}

// Fields of the PulseValue.
func (PulseValue) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint("timestamp"),
		field.String("type"),
		field.String("key"),
		field.Bytes("key_hash").Optional().Nillable(),
		field.String("value"),
	}
}

// Edges of the PulseValue.
func (PulseValue) Edges() []ent.Edge {
	return []ent.Edge{}
}
