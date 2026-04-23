package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AdminInvite holds the schema definition for the AdminInvite entity.
type AdminInvite struct {
	ent.Schema
}

// Annotations of the AdminInvite.
func (AdminInvite) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "admin_invites"},
	}
}

// Fields of the AdminInvite.
func (AdminInvite) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Optional().Nillable(),
		field.String("invite_code").Unique(),
		field.String("description").Optional().Nillable(),
		field.String("message").Optional().Nillable(),
		field.Uint("max_uses").Optional().Nillable(),
		field.Uint("uses"),
		field.Bool("skip_email_verification"),
		field.Time("expires_at").Optional().Nillable(),
		field.String("used_by").Optional().Nillable(),
		field.Uint("admin_user_id").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the AdminInvite.
func (AdminInvite) Edges() []ent.Edge {
	return []ent.Edge{}
}
