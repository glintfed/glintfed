package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupLike holds the schema definition for the GroupLike entity.
type GroupLike struct {
	ent.Schema
}

// Annotations of the GroupLike.
func (GroupLike) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_likes"},
	}
}

// Fields of the GroupLike.
func (GroupLike) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.Uint64("status_id").Optional().Nillable(),
		field.Uint64("comment_id").Optional().Nillable(),
		field.Bool("local"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupLike.
func (GroupLike) Edges() []ent.Edge {
	return []ent.Edge{}
}
