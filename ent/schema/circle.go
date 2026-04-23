package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Circle holds the schema definition for the Circle entity.
type Circle struct {
	ent.Schema
}

// Annotations of the Circle.
func (Circle) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "circles"},
	}
}

// Fields of the Circle.
func (Circle) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.String("name").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.String("scope"),
		field.Bool("bcc"),
		field.Bool("active"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Circle.
func (Circle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Profile.Type).Ref("circles").Field("profile_id").Unique().Required(),
	}
}
