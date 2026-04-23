package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Annotations of the Job.
func (Job) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "jobs"},
	}
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("queue"),
		field.String("payload"),
		field.Uint("attempts"),
		field.Uint("reserved_at").Optional().Nillable(),
		field.Uint("available_at"),
		field.Uint("created_at"),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{}
}
