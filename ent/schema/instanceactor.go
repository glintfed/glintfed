package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// InstanceActor holds the schema definition for the InstanceActor entity.
type InstanceActor struct {
	ent.Schema
}

// Annotations of the InstanceActor.
func (InstanceActor) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "instance_actors"},
	}
}

// Fields of the InstanceActor.
func (InstanceActor) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("private_key").Optional().Nillable(),
		field.String("public_key").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the InstanceActor.
func (InstanceActor) Edges() []ent.Edge {
	return []ent.Edge{}
}
