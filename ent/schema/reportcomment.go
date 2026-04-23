package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ReportComment holds the schema definition for the ReportComment entity.
type ReportComment struct {
	ent.Schema
}

// Annotations of the ReportComment.
func (ReportComment) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "report_comments"},
	}
}

// Fields of the ReportComment.
func (ReportComment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.Uint64("report_id"),
		field.Uint64("profile_id"),
		field.Uint64("user_id"),
		field.String("comment"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ReportComment.
func (ReportComment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
	}
}
