package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CuratedRegister holds the schema definition for the CuratedRegister entity.
type CuratedRegister struct {
	ent.Schema
}

// Annotations of the CuratedRegister.
func (CuratedRegister) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "curated_registers"},
	}
}

// Fields of the CuratedRegister.
func (CuratedRegister) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("email").Optional().Nillable().Unique(),
		field.String("username").Optional().Nillable().Unique(),
		field.String("password").Optional().Nillable(),
		field.String("ip_address").Optional().Nillable(),
		field.String("verify_code").Optional().Nillable(),
		field.String("reason_to_join").Optional().Nillable(),
		field.Uint64("invited_by").Optional().Nillable(),
		field.Bool("is_approved"),
		field.Bool("is_rejected"),
		field.Bool("is_awaiting_more_info"),
		field.Bool("user_has_responded"),
		field.Bool("is_closed"),
		field.String("autofollow_account_ids").Optional().Nillable(),
		field.String("admin_notes").Optional().Nillable(),
		field.Uint("approved_by_admin_id").Optional().Nillable(),
		field.Time("email_verified_at").Optional().Nillable(),
		field.Time("admin_notified_at").Optional().Nillable(),
		field.Time("action_taken_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CuratedRegister.
func (CuratedRegister) Edges() []ent.Edge {
	return []ent.Edge{}
}
