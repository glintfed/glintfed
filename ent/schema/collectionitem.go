package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CollectionItem holds the schema definition for the CollectionItem entity.
type CollectionItem struct {
	ent.Schema
}

// Annotations of the CollectionItem.
func (CollectionItem) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "collection_items"},
	}
}

// Fields of the CollectionItem.
func (CollectionItem) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("collection_id"),
		field.Uint("order").Optional().Nillable(),
		field.String("object_type"),
		field.Uint64("object_id"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the CollectionItem.
func (CollectionItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("collection", Collection.Type).Ref("items").Field("collection_id").Unique().Required(),
	}
}
