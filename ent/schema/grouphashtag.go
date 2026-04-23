package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupHashtag holds the schema definition for the GroupHashtag entity.
type GroupHashtag struct {
	ent.Schema
}

// Annotations of the GroupHashtag.
func (GroupHashtag) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "group_hashtags"},
	}
}

// Fields of the GroupHashtag.
func (GroupHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Unique(),
		field.String("formatted").Optional().Nillable(),
		field.Bool("recommended"),
		field.Bool("sensitive"),
		field.Bool("banned"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the GroupHashtag.
func (GroupHashtag) Edges() []ent.Edge {
	return []ent.Edge{}
}
