package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ImportJob holds the schema definition for the ImportJob entity.
type ImportJob struct {
	ent.Schema
}

// Annotations of the ImportJob.
func (ImportJob) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "import_jobs"},
	}
}

// Fields of the ImportJob.
func (ImportJob) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.Uint64("profile_id"),
		field.String("service"),
		field.String("uuid").Optional().Nillable(),
		field.String("storage_path").Optional().Nillable(),
		field.Uint("stage"),
		field.String("media_json").Optional().Nillable(),
		field.Time("completed_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ImportJob.
func (ImportJob) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
		edge.To("files", ImportData.Type),
	}
}
