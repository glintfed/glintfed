package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// LiveStream holds the schema definition for the LiveStream entity.
type LiveStream struct {
	ent.Schema
}

// Annotations of the LiveStream.
func (LiveStream) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "live_streams"},
	}
}

// Fields of the LiveStream.
func (LiveStream) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.String("stream_id").Optional().Nillable().Unique(),
		field.String("stream_key").Optional().Nillable(),
		field.String("visibility").Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.String("thumbnail_path").Optional().Nillable(),
		field.String("settings").Optional().Nillable(),
		field.Bool("live_chat"),
		field.String("mod_ids").Optional().Nillable(),
		field.Bool("discoverable").Optional().Nillable(),
		field.Time("live_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the LiveStream.
func (LiveStream) Edges() []ent.Edge {
	return []ent.Edge{}
}
