package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CustomFilter holds the schema definition for the CustomFilter entity.
type CustomFilter struct {
	ent.Schema
}

// Annotations of the CustomFilter.
func (CustomFilter) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "custom_filters"},
	}
}

// Fields of the CustomFilter.
func (CustomFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.String("phrase"),
		field.Int("action"),
		field.String("context").Optional().Nillable(),
		field.Time("expires_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CustomFilter.
func (CustomFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Profile.Type).Field("profile_id").Unique().Required(),
		edge.To("keywords", CustomFilterKeyword.Type),
		edge.To("statuses", CustomFilterStatus.Type),
	}
}
