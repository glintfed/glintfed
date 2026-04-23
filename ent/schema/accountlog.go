package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccountLog holds the schema definition for the AccountLog entity.
type AccountLog struct {
	ent.Schema
}

// Annotations of the AccountLog.
func (AccountLog) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "account_logs"},
	}
}

// Fields of the AccountLog.
func (AccountLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id"),
		field.Uint64("item_id").Optional().Nillable(),
		field.String("item_type").Optional().Nillable(),
		field.String("action").Optional().Nillable(),
		field.String("message").Optional().Nillable(),
		field.String("link").Optional().Nillable(),
		field.String("ip_address").Optional().Nillable(),
		field.String("user_agent").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the AccountLog.
func (AccountLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("accountLog").Field("user_id").Unique().Required(),
	}
}
