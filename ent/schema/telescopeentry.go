package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TelescopeEntry holds the schema definition for the TelescopeEntry entity.
type TelescopeEntry struct {
	ent.Schema
}

// Annotations of the TelescopeEntry.
func (TelescopeEntry) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "telescope_entries"},
	}
}

// Fields of the TelescopeEntry.
func (TelescopeEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("sequence"),
		field.String("uuid").Unique(),
		field.String("batch_id"),
		field.String("family_hash").Optional().Nillable(),
		field.Bool("should_display_on_index"),
		field.String("type"),
		field.String("content"),
		field.Time("created_at").Optional().Nillable(),
	}
}

// Edges of the TelescopeEntry.
func (TelescopeEntry) Edges() []ent.Edge {
	return []ent.Edge{}
}
