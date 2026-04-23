package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Annotations of the Like.
func (Like) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "likes"},
	}
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.Uint64("status_id"),
		field.Uint64("status_profile_id").Optional().Nillable(),
		field.Bool("is_comment").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("actor", Profile.Type).Ref("likes").Field("profile_id").Unique().Required(),
		edge.From("status", Status.Type).Ref("likes").Field("status_id").Unique().Required(),
	}
}
