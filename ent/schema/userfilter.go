package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserFilter holds the schema definition for the UserFilter entity.
type UserFilter struct {
	ent.Schema
}

// Annotations of the UserFilter.
func (UserFilter) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_filters"},
	}
}

// Fields of the UserFilter.
func (UserFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id"),
		field.Uint64("filterable_id"),
		field.String("filterable_type"),
		field.String("filter_type"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the UserFilter.
func (UserFilter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("instance", Instance.Type).Field("filterable_id").Unique().Required(),
		edge.To("user", Profile.Type).Field("user_id").Unique().Required(),
	}
}
