package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProfileMigration holds the schema definition for the ProfileMigration entity.
type ProfileMigration struct {
	ent.Schema
}

// Annotations of the ProfileMigration.
func (ProfileMigration) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "profile_migrations"},
	}
}

// Fields of the ProfileMigration.
func (ProfileMigration) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.String("acct").Optional().Nillable(),
		field.Uint64("followers_count"),
		field.Uint64("target_profile_id").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ProfileMigration.
func (ProfileMigration) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
		edge.To("target", Profile.Type).Field("target_profile_id").Unique(),
	}
}
