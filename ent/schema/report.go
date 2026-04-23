package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Annotations of the Report.
func (Report) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "reports"},
	}
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id"),
		field.Uint64("user_id").Optional().Nillable(),
		field.Uint64("object_id"),
		field.String("object_type").Optional().Nillable(),
		field.Uint64("reported_profile_id").Optional().Nillable(),
		field.String("type").Optional().Nillable(),
		field.String("message").Optional().Nillable(),
		field.Time("admin_seen").Optional().Nillable(),
		field.Bool("not_interested"),
		field.Bool("spam"),
		field.Bool("nsfw"),
		field.Bool("abusive"),
		field.String("meta").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("reporter", Profile.Type).Ref("reports").Field("profile_id").Unique().Required(),
		edge.To("status", Status.Type).Field("object_id").Unique().Required(),
		edge.To("reportedUser", Profile.Type).Field("reported_profile_id").Unique(),
	}
}
