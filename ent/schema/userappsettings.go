package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserAppSettings holds the schema definition for the UserAppSettings entity.
type UserAppSettings struct {
	ent.Schema
}

// Annotations of the UserAppSettings.
func (UserAppSettings) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_app_settings"},
	}
}

// Fields of the UserAppSettings.
func (UserAppSettings) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id").Unique(),
		field.Uint64("profile_id").Unique(),
		field.String("common").Optional().Nillable(),
		field.String("custom").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the UserAppSettings.
func (UserAppSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique().Required(),
	}
}
