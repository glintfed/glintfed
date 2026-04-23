package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PulseEntry holds the schema definition for the PulseEntry entity.
type PulseEntry struct {
	ent.Schema
}

// Annotations of the PulseEntry.
func (PulseEntry) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "pulse_entries"},
	}
}

// Fields of the PulseEntry.
func (PulseEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint("timestamp"),
		field.String("type"),
		field.String("key"),
		field.Bytes("key_hash").Optional().Nillable(),
		field.Int64("value").Optional().Nillable(),
	}
}

// Edges of the PulseEntry.
func (PulseEntry) Edges() []ent.Edge {
	return []ent.Edge{}
}
