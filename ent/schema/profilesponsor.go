package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProfileSponsor holds the schema definition for the ProfileSponsor entity.
type ProfileSponsor struct {
	ent.Schema
}

// Annotations of the ProfileSponsor.
func (ProfileSponsor) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "profile_sponsors"},
	}
}

// Fields of the ProfileSponsor.
func (ProfileSponsor) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id").Unique(),
		field.String("sponsors").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ProfileSponsor.
func (ProfileSponsor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
	}
}
