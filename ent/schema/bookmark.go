package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bookmark holds the schema definition for the Bookmark entity.
type Bookmark struct {
	ent.Schema
}

// Annotations of the Bookmark.
func (Bookmark) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "bookmarks"},
	}
}

// Fields of the Bookmark.
func (Bookmark) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.Uint64("status_id"),
		field.Uint64("profile_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status", Status.Type).Field("status_id").Unique().Required(),
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
	}
}
