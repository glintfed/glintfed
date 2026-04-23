package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CuratedRegisterActivity holds the schema definition for the CuratedRegisterActivity entity.
type CuratedRegisterActivity struct {
	ent.Schema
}

// Annotations of the CuratedRegisterActivity.
func (CuratedRegisterActivity) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "curated_register_activities"},
	}
}

// Fields of the CuratedRegisterActivity.
func (CuratedRegisterActivity) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("register_id").Optional().Nillable(),
		field.Uint("admin_id").Optional().Nillable(),
		field.Uint("reply_to_id").Optional().Nillable(),
		field.String("secret_code").Optional().Nillable(),
		field.String("type").Optional().Nillable(),
		field.String("title").Optional().Nillable(),
		field.String("link").Optional().Nillable(),
		field.String("message").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Bool("from_admin"),
		field.Bool("from_user"),
		field.Bool("admin_only_view"),
		field.Bool("action_required"),
		field.Time("admin_notified_at").Optional().Nillable(),
		field.Time("action_taken_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CuratedRegisterActivity.
func (CuratedRegisterActivity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("application", CuratedRegister.Type).Field("register_id").Unique(),
	}
}
