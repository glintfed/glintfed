package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupPostHashtag holds the schema definition for the GroupPostHashtag entity.
type GroupPostHashtag struct {
	ent.Schema
}

// Annotations of the GroupPostHashtag.
func (GroupPostHashtag) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_post_hashtags"},
	}
}

// Fields of the GroupPostHashtag.
func (GroupPostHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("hashtag_id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.Uint64("status_id").Optional().Nillable(),
		field.String("status_visibility").Optional().Nillable(),
		field.Bool("nsfw"),
	}
}

// Edges of the GroupPostHashtag.
func (GroupPostHashtag) Edges() []ent.Edge {
	return []ent.Edge{}
}
