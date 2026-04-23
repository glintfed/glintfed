package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupBlock holds the schema definition for the GroupBlock entity.
type GroupBlock struct {
	ent.Schema
}

// Annotations of the GroupBlock.
func (GroupBlock) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_blocks"},
	}
}

// Fields of the GroupBlock.
func (GroupBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("group_id"),
		field.Uint64("admin_id").Optional().Nillable(),
		field.Uint64("profile_id").Optional().Nillable(),
		field.Uint64("instance_id").Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("reason").Optional().Nillable(),
		field.Bool("is_user"),
		field.Bool("moderated"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupBlock.
func (GroupBlock) Edges() []ent.Edge {
	return []ent.Edge{}
}
