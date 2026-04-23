package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// JobBatch holds the schema definition for the JobBatch entity.
type JobBatch struct {
	ent.Schema
}

// Annotations of the JobBatch.
func (JobBatch) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "job_batches"},
	}
}

// Fields of the JobBatch.
func (JobBatch) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.Int("total_jobs"),
		field.Int("pending_jobs"),
		field.Int("failed_jobs"),
		field.String("failed_job_ids"),
		field.String("options").Optional().Nillable(),
		field.Int("cancelled_at").Optional().Nillable(),
		field.Int("created_at"),
		field.Int("finished_at").Optional().Nillable(),
	}
}

// Edges of the JobBatch.
func (JobBatch) Edges() []ent.Edge {
	return []ent.Edge{}
}
