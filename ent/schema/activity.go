package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Annotations of the Activity.
func (Activity) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "activities"},
	}
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("to_id").Optional().Nillable(),
		field.Uint64("from_id").Optional().Nillable(),
		field.String("object_type").Optional().Nillable(),
		field.String("data").Optional().Nillable(),
		field.Time("processed_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("toProfile", Profile.Type).Field("to_id").Unique(),
		edge.To("fromProfile", Profile.Type).Field("from_id").Unique(),
	}
}
