package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CircleProfile holds the schema definition for the CircleProfile entity.
type CircleProfile struct {
	ent.Schema
}

// Annotations of the CircleProfile.
func (CircleProfile) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "circle_profiles"},
	}
}

// Fields of the CircleProfile.
func (CircleProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("owner_id").Optional().Nillable(),
		field.Uint64("circle_id"),
		field.Uint64("profile_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CircleProfile.
func (CircleProfile) Edges() []ent.Edge {
	return []ent.Edge{}
}
