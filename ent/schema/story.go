package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Story holds the schema definition for the Story entity.
type Story struct {
	ent.Schema
}

// Annotations of the Story.
func (Story) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "stories"},
	}
}

// Fields of the Story.
func (Story) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.String("type").Optional().Nillable(),
		field.Uint("size").Optional().Nillable(),
		field.String("mime").Optional().Nillable(),
		field.Uint("duration"),
		field.String("path").Optional().Nillable(),
		field.String("remote_url").Optional().Nillable().Unique(),
		field.String("media_url").Optional().Nillable().Unique(),
		field.String("cdn_url").Optional().Nillable(),
		field.Bool("public"),
		field.Bool("local"),
		field.Uint("view_count"),
		field.Uint("comment_count").Optional().Nillable(),
		field.String("story").Optional().Nillable(),
		field.Time("expires_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Bool("is_archived").Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.Bool("active").Optional().Nillable(),
		field.Bool("can_reply"),
		field.Bool("can_react"),
		field.String("object_id").Optional().Nillable().Unique(),
		field.String("object_uri").Optional().Nillable().Unique(),
		field.String("bearcap_token").Optional().Nillable(),
	}
}

// Edges of the Story.
func (Story) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("stories").Field("profile_id").Unique().Required(),
		edge.To("views", StoryView.Type),
	}
}
