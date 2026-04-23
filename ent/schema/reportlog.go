package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ReportLog holds the schema definition for the ReportLog entity.
type ReportLog struct {
	ent.Schema
}

// Annotations of the ReportLog.
func (ReportLog) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "report_logs"},
	}
}

// Fields of the ReportLog.
func (ReportLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.Uint64("profile_id"),
		field.Uint64("item_id").Optional().Nillable(),
		field.String("item_type").Optional().Nillable(),
		field.String("action").Optional().Nillable(),
		field.Bool("system_message"),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ReportLog.
func (ReportLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
	}
}
