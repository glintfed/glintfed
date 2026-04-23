package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

// Annotations of the Portfolio.
func (Portfolio) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "portfolios"},
	}
}

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint("user_id").Optional().Nillable().Unique(),
		field.Uint64("profile_id").Unique(),
		field.Bool("active").Optional().Nillable(),
		field.Bool("show_captions").Optional().Nillable(),
		field.Bool("show_license").Optional().Nillable(),
		field.Bool("show_location").Optional().Nillable(),
		field.Bool("show_timestamp").Optional().Nillable(),
		field.Bool("show_link").Optional().Nillable(),
		field.String("profile_source").Optional().Nillable(),
		field.Bool("show_avatar").Optional().Nillable(),
		field.Bool("show_bio").Optional().Nillable(),
		field.String("profile_layout").Optional().Nillable(),
		field.String("profile_container").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return []ent.Edge{}
}
