package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Annotations of the Collection.
func (Collection) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "collections"},
	}
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id").Optional().Nillable(),
		field.String("title").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.Bool("is_nsfw"),
		field.String("visibility"),
		field.Time("published_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("collections").Field("profile_id").Unique(),
		edge.To("items", CollectionItem.Type),
	}
}
