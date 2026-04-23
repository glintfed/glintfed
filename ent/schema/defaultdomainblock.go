package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// DefaultDomainBlock holds the schema definition for the DefaultDomainBlock entity.
type DefaultDomainBlock struct {
	ent.Schema
}

// Annotations of the DefaultDomainBlock.
func (DefaultDomainBlock) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "default_domain_blocks"},
	}
}

// Fields of the DefaultDomainBlock.
func (DefaultDomainBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("domain").Unique(),
		field.String("note").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the DefaultDomainBlock.
func (DefaultDomainBlock) Edges() []ent.Edge {
	return []ent.Edge{}
}
