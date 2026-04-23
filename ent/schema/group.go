package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Annotations of the Group.
func (Group) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "groups"},
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint("category_id"),
		field.Uint64("profile_id").Optional().Nillable(),
		field.String("status").Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.String("rules").Optional().Nillable(),
		field.Bool("local"),
		field.String("remote_url").Optional().Nillable(),
		field.String("inbox_url").Optional().Nillable(),
		field.Bool("is_private"),
		field.Bool("local_only"),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Uint("member_count").Optional().Nillable(),
		field.Bool("recommended"),
		field.Bool("discoverable"),
		field.Bool("activitypub"),
		field.Bool("is_nsfw"),
		field.Bool("dms"),
		field.Bool("autospam"),
		field.Bool("verified"),
		field.Time("last_active_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", GroupMember.Type),
		edge.To("admin", Profile.Type).Field("profile_id").Unique(),
	}
}
