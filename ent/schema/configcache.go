package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ConfigCache holds the schema definition for the ConfigCache entity.
type ConfigCache struct {
	ent.Schema
}

// Annotations of the ConfigCache.
func (ConfigCache) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "config_cache"},
	}
}

// Fields of the ConfigCache.
func (ConfigCache) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("k").Unique(),
		field.String("v").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ConfigCache.
func (ConfigCache) Edges() []ent.Edge {
	return []ent.Edge{}
}
