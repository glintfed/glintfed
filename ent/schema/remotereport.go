package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RemoteReport holds the schema definition for the RemoteReport entity.
type RemoteReport struct {
	ent.Schema
}

// Annotations of the RemoteReport.
func (RemoteReport) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "remote_reports"},
	}
}

// Fields of the RemoteReport.
func (RemoteReport) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("status_ids").Optional().Nillable(),
		field.String("comment").Optional().Nillable(),
		field.Uint64("account_id").Optional().Nillable(),
		field.String("uri").Optional().Nillable(),
		field.Uint("instance_id").Optional().Nillable(),
		field.Time("action_taken_at").Optional().Nillable(),
		field.String("report_meta").Optional().Nillable(),
		field.String("action_taken_meta").Optional().Nillable(),
		field.Uint64("action_taken_by_account_id").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the RemoteReport.
func (RemoteReport) Edges() []ent.Edge {
	return []ent.Edge{}
}
