package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AppRegister holds the schema definition for the AppRegister entity.
type AppRegister struct {
	ent.Schema
}

// Annotations of the AppRegister.
func (AppRegister) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "app_registers"},
	}
}

// Fields of the AppRegister.
func (AppRegister) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("email"),
		field.String("verify_code"),
		field.Time("email_delivered_at").Optional().Nillable(),
		field.Time("email_verified_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Uint("uses"),
	}
}

// Edges of the AppRegister.
func (AppRegister) Edges() []ent.Edge {
	return []ent.Edge{}
}
