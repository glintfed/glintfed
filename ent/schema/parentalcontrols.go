package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ParentalControls holds the schema definition for the ParentalControls entity.
type ParentalControls struct {
	ent.Schema
}

// Annotations of the ParentalControls.
func (ParentalControls) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "parental_controls"},
	}
}

// Fields of the ParentalControls.
func (ParentalControls) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("parent_id"),
		field.Uint64("child_id").Optional().Nillable().Unique(),
		field.String("email").Optional().Nillable().Unique(),
		field.String("verify_code").Optional().Nillable(),
		field.Time("email_sent_at").Optional().Nillable(),
		field.Time("email_verified_at").Optional().Nillable(),
		field.String("permissions").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ParentalControls.
func (ParentalControls) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", User.Type).Field("parent_id").Unique().Required(),
		edge.To("child", User.Type).Field("child_id").Unique(),
	}
}
