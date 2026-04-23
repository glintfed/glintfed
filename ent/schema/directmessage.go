package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DirectMessage holds the schema definition for the DirectMessage entity.
type DirectMessage struct {
	ent.Schema
}

// Annotations of the DirectMessage.
func (DirectMessage) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "direct_messages"},
	}
}

// Fields of the DirectMessage.
func (DirectMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("to_id"),
		field.Uint64("from_id"),
		field.String("type").Optional().Nillable(),
		field.String("from_profile_ids").Optional().Nillable(),
		field.Bool("group_message"),
		field.Bool("is_hidden"),
		field.String("meta").Optional().Nillable(),
		field.Uint64("status_id"),
		field.Time("read_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the DirectMessage.
func (DirectMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("status", Status.Type).Ref("directMessage").Field("status_id").Unique().Required(),
		edge.To("author", Profile.Type).Field("from_id").Unique().Required(),
		edge.To("recipient", Profile.Type).Field("to_id").Unique().Required(),
	}
}
