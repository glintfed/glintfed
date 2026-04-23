package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupMember holds the schema definition for the GroupMember entity.
type GroupMember struct {
	ent.Schema
}

// Annotations of the GroupMember.
func (GroupMember) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_members"},
	}
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.String("role"),
		field.Bool("local_group"),
		field.Bool("local_profile"),
		field.Bool("join_request"),
		field.Time("approved_at").Optional().Nillable(),
		field.Time("rejected_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("members").Field("group_id").Unique().Required(),
	}
}
