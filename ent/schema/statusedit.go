package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StatusEdit holds the schema definition for the StatusEdit entity.
type StatusEdit struct {
	ent.Schema
}

// Annotations of the StatusEdit.
func (StatusEdit) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "status_edits"},
	}
}

// Fields of the StatusEdit.
func (StatusEdit) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("status_id"),
		field.Uint64("profile_id"),
		field.String("caption").Optional().Nillable(),
		field.String("spoiler_text").Optional().Nillable(),
		field.String("ordered_media_attachment_ids").Optional().Nillable(),
		field.String("media_descriptions").Optional().Nillable(),
		field.String("poll_options").Optional().Nillable(),
		field.Bool("is_nsfw").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the StatusEdit.
func (StatusEdit) Edges() []ent.Edge {
	return []ent.Edge{}
}
