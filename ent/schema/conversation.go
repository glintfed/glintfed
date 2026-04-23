package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Conversation holds the schema definition for the Conversation entity.
type Conversation struct {
	ent.Schema
}

// Annotations of the Conversation.
func (Conversation) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "conversations"},
	}
}

// Fields of the Conversation.
func (Conversation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("to_id"),
		field.Uint64("from_id"),
		field.Uint64("dm_id").Optional().Nillable(),
		field.Uint64("status_id").Optional().Nillable(),
		field.String("type").Optional().Nillable(),
		field.Bool("is_hidden"),
		field.Bool("has_seen"),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Conversation.
func (Conversation) Edges() []ent.Edge {
	return []ent.Edge{}
}
