package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ModLog holds the schema definition for the ModLog entity.
type ModLog struct {
	ent.Schema
}

// Annotations of the ModLog.
func (ModLog) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "mod_logs"},
	}
}

// Fields of the ModLog.
func (ModLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id"),
		field.String("user_username").Optional().Nillable(),
		field.Uint64("object_uid").Optional().Nillable(),
		field.Uint64("object_id").Optional().Nillable(),
		field.String("object_type").Optional().Nillable(),
		field.String("action").Optional().Nillable(),
		field.String("message").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.String("access_level").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ModLog.
func (ModLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admin", User.Type).Field("user_id").Unique().Required(),
	}
}
