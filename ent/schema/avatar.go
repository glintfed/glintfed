package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Avatar holds the schema definition for the Avatar entity.
type Avatar struct {
	ent.Schema
}

// Annotations of the Avatar.
func (Avatar) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "avatars"},
	}
}

// Fields of the Avatar.
func (Avatar) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id").Unique(),
		field.String("media_path").Optional().Nillable(),
		field.String("remote_url").Optional().Nillable(),
		field.String("cdn_url").Optional().Nillable(),
		field.Bool("is_remote").Optional().Nillable(),
		field.Uint("size").Optional().Nillable(),
		field.Uint("change_count"),
		field.Time("last_fetched_at").Optional().Nillable(),
		field.Time("last_processed_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Avatar.
func (Avatar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("avatar").Field("profile_id").Unique().Required(),
	}
}
