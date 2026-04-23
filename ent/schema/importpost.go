package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ImportPost holds the schema definition for the ImportPost entity.
type ImportPost struct {
	ent.Schema
}

// Annotations of the ImportPost.
func (ImportPost) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "import_posts"},
	}
}

// Fields of the ImportPost.
func (ImportPost) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.Uint("user_id"),
		field.String("service"),
		field.String("post_hash").Optional().Nillable(),
		field.String("filename"),
		field.Uint("media_count"),
		field.String("post_type").Optional().Nillable(),
		field.String("caption").Optional().Nillable(),
		field.String("media").Optional().Nillable(),
		field.Uint("creation_year").Optional().Nillable(),
		field.Uint("creation_month").Optional().Nillable(),
		field.Uint("creation_day").Optional().Nillable(),
		field.Uint("creation_id").Optional().Nillable(),
		field.Uint64("status_id").Optional().Nillable().Unique(),
		field.Time("creation_date").Optional().Nillable(),
		field.String("metadata").Optional().Nillable(),
		field.Bool("skip_missing_media"),
		field.Bool("uploaded_to_s3"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the ImportPost.
func (ImportPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("status", Status.Type).Field("status_id").Unique(),
	}
}
