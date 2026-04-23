package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// HashtagFollow holds the schema definition for the HashtagFollow entity.
type HashtagFollow struct {
	ent.Schema
}

// Annotations of the HashtagFollow.
func (HashtagFollow) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "hashtag_follows"},
	}
}

// Fields of the HashtagFollow.
func (HashtagFollow) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id"),
		field.Uint64("profile_id"),
		field.Uint64("hashtag_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the HashtagFollow.
func (HashtagFollow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hashtag", Hashtag.Type).Field("hashtag_id").Unique().Required(),
	}
}
