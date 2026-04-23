package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupStore holds the schema definition for the GroupStore entity.
type GroupStore struct {
	ent.Schema
}

// Annotations of the GroupStore.
func (GroupStore) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_stores"},
	}
}

// Fields of the GroupStore.
func (GroupStore) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id").Optional().Nillable(),
		field.String("store_key"),
		field.String("store_value").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupStore.
func (GroupStore) Edges() []ent.Edge {
	return []ent.Edge{}
}
