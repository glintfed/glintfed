package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TelescopeMonitoring holds the schema definition for the TelescopeMonitoring entity.
type TelescopeMonitoring struct {
	ent.Schema
}

// Annotations of the TelescopeMonitoring.
func (TelescopeMonitoring) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "telescope_monitoring"},
	}
}

// Fields of the TelescopeMonitoring.
func (TelescopeMonitoring) Fields() []ent.Field {
	return []ent.Field{
		field.String("tag"),
	}
}

// Edges of the TelescopeMonitoring.
func (TelescopeMonitoring) Edges() []ent.Edge {
	return []ent.Edge{}
}
