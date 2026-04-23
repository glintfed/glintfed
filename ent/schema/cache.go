package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Cache holds the schema definition for the Cache entity.
type Cache struct {
	ent.Schema
}

// Annotations of the Cache.
func (Cache) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "cache"},
	}
}

// Fields of the Cache.
func (Cache) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.String("value"),
		field.Int("expiration"),
	}
}

// Edges of the Cache.
func (Cache) Edges() []ent.Edge {
	return []ent.Edge{}
}
