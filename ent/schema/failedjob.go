package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// FailedJob holds the schema definition for the FailedJob entity.
type FailedJob struct {
	ent.Schema
}

// Annotations of the FailedJob.
func (FailedJob) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "failed_jobs"},
	}
}

// Fields of the FailedJob.
func (FailedJob) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("uuid").Optional().Nillable().Unique(),
		field.String("connection"),
		field.String("queue"),
		field.String("payload"),
		field.String("exception"),
		field.Time("failed_at"),
	}
}

// Edges of the FailedJob.
func (FailedJob) Edges() []ent.Edge {
	return []ent.Edge{}
}
