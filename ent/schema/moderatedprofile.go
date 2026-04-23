package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ModeratedProfile holds the schema definition for the ModeratedProfile entity.
type ModeratedProfile struct {
	ent.Schema
}

// Annotations of the ModeratedProfile.
func (ModeratedProfile) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "moderated_profiles"},
	}
}

// Fields of the ModeratedProfile.
func (ModeratedProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("profile_url").Optional().Nillable().Unique(),
		field.Uint64("profile_id").Optional().Nillable().Unique(),
		field.String("domain").Optional().Nillable(),
		field.String("note").Optional().Nillable(),
		field.Bool("is_banned"),
		field.Bool("is_nsfw"),
		field.Bool("is_unlisted"),
		field.Bool("is_noautolink"),
		field.Bool("is_nodms"),
		field.Bool("is_notrending"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ModeratedProfile.
func (ModeratedProfile) Edges() []ent.Edge {
	return []ent.Edge{}
}
