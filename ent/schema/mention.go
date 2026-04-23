package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mention holds the schema definition for the Mention entity.
type Mention struct {
	ent.Schema
}

// Annotations of the Mention.
func (Mention) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "mentions"},
	}
}

// Fields of the Mention.
func (Mention) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("status_id"),
		field.Uint64("profile_id"),
		field.Bool("local"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Mention.
func (Mention) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
		edge.To("status", Status.Type).Field("status_id").Unique().Required(),
	}
}
