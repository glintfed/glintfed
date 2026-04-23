package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TelescopeEntriesTag holds the schema definition for the TelescopeEntriesTag entity.
type TelescopeEntriesTag struct {
	ent.Schema
}

// Annotations of the TelescopeEntriesTag.
func (TelescopeEntriesTag) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "telescope_entries_tags"},
	}
}

// Fields of the TelescopeEntriesTag.
func (TelescopeEntriesTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("entry_uuid"),
		field.String("tag"),
	}
}

// Edges of the TelescopeEntriesTag.
func (TelescopeEntriesTag) Edges() []ent.Edge {
	return []ent.Edge{}
}
