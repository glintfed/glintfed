package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CacheLock holds the schema definition for the CacheLock entity.
type CacheLock struct {
	ent.Schema
}

// Annotations of the CacheLock.
func (CacheLock) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "cache_locks"},
	}
}

// Fields of the CacheLock.
func (CacheLock) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.String("owner"),
		field.Int("expiration"),
	}
}

// Edges of the CacheLock.
func (CacheLock) Edges() []ent.Edge {
	return []ent.Edge{}
}
