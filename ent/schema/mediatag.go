package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MediaTag holds the schema definition for the MediaTag entity.
type MediaTag struct {
	ent.Schema
}

// Annotations of the MediaTag.
func (MediaTag) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "media_tags"},
	}
}

// Fields of the MediaTag.
func (MediaTag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("status_id").Optional().Nillable(),
		field.Uint64("media_id"),
		field.Uint64("profile_id"),
		field.String("tagged_username").Optional().Nillable(),
		field.Bool("is_public"),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the MediaTag.
func (MediaTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status", Status.Type).Field("status_id").Unique(),
	}
}
