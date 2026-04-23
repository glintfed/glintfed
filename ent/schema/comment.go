package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Annotations of the Comment.
func (Comment) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "comments"},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.Uint64("user_id").Optional().Nillable(),
		field.Uint64("status_id"),
		field.String("comment").Optional().Nillable(),
		field.String("rendered").Optional().Nillable(),
		field.String("entities").Optional().Nillable(),
		field.Bool("is_remote"),
		field.Time("rendered_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
		edge.To("status", Status.Type).Field("status_id").Unique().Required(),
	}
}
