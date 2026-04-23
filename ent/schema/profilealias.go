package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProfileAlias holds the schema definition for the ProfileAlias entity.
type ProfileAlias struct {
	ent.Schema
}

// Annotations of the ProfileAlias.
func (ProfileAlias) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "profile_aliases"},
	}
}

// Fields of the ProfileAlias.
func (ProfileAlias) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id").Optional().Nillable(),
		field.String("acct").Optional().Nillable(),
		field.String("uri").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ProfileAlias.
func (ProfileAlias) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("aliases").Field("profile_id").Unique(),
	}
}
