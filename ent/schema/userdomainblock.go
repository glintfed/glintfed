package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserDomainBlock holds the schema definition for the UserDomainBlock entity.
type UserDomainBlock struct {
	ent.Schema
}

// Annotations of the UserDomainBlock.
func (UserDomainBlock) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "user_domain_blocks"},
	}
}

// Fields of the UserDomainBlock.
func (UserDomainBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.String("domain"),
	}
}

// Edges of the UserDomainBlock.
func (UserDomainBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Field("profile_id").Unique().Required(),
	}
}
