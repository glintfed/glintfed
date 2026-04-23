package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StatusHashtag holds the schema definition for the StatusHashtag entity.
type StatusHashtag struct {
	ent.Schema
}

// Annotations of the StatusHashtag.
func (StatusHashtag) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "status_hashtags"},
	}
}

// Fields of the StatusHashtag.
func (StatusHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("status_id"),
		field.Uint64("hashtag_id"),
		field.Uint64("profile_id").Optional().Nillable(),
		field.String("status_visibility").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the StatusHashtag.
func (StatusHashtag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status", Status.Type).Field("status_id").Unique().Required(),
		edge.To("hashtag", Hashtag.Type).Field("hashtag_id").Unique().Required(),
		edge.To("profile", Profile.Type).Field("profile_id").Unique(),
	}
}
