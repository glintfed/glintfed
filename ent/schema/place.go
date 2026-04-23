package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Place holds the schema definition for the Place entity.
type Place struct {
	ent.Schema
}

// Annotations of the Place.
func (Place) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "places"},
	}
}

// Fields of the Place.
func (Place) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("slug"),
		field.String("name"),
		field.String("state").Optional().Nillable(),
		field.String("country"),
		field.String("aliases").Optional().Nillable(),
		field.Float("lat").Optional().Nillable(),
		field.Float("long").Optional().Nillable(),
		field.Int("score"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Uint64("cached_post_count").Optional().Nillable(),
		field.Time("last_checked_at").Optional().Nillable(),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Status.Type),
		edge.To("statuses", Status.Type),
	}
}
