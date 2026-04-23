package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MediaBlocklist holds the schema definition for the MediaBlocklist entity.
type MediaBlocklist struct {
	ent.Schema
}

// Annotations of the MediaBlocklist.
func (MediaBlocklist) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "media_blocklists"},
	}
}

// Fields of the MediaBlocklist.
func (MediaBlocklist) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("sha256").Optional().Nillable().Unique(),
		field.String("sha512").Optional().Nillable().Unique(),
		field.String("name").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.Bool("active"),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the MediaBlocklist.
func (MediaBlocklist) Edges() []ent.Edge {
	return []ent.Edge{}
}
