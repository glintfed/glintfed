package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AutospamCustomTokens holds the schema definition for the AutospamCustomTokens entity.
type AutospamCustomTokens struct {
	ent.Schema
}

// Annotations of the AutospamCustomTokens.
func (AutospamCustomTokens) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "autospam_custom_tokens"},
	}
}

// Fields of the AutospamCustomTokens.
func (AutospamCustomTokens) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("token"),
		field.Int("weight"),
		field.Bool("is_spam"),
		field.String("note").Optional().Nillable(),
		field.String("category").Optional().Nillable(),
		field.Bool("active"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the AutospamCustomTokens.
func (AutospamCustomTokens) Edges() []ent.Edge {
	return []ent.Edge{}
}
