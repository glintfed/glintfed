package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccountInterstitial holds the schema definition for the AccountInterstitial entity.
type AccountInterstitial struct {
	ent.Schema
}

// Annotations of the AccountInterstitial.
func (AccountInterstitial) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "account_interstitials"},
	}
}

// Fields of the AccountInterstitial.
func (AccountInterstitial) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("user_id").Optional().Nillable(),
		field.String("type").Optional().Nillable(),
		field.String("view").Optional().Nillable(),
		field.Uint64("item_id").Optional().Nillable(),
		field.String("item_type").Optional().Nillable(),
		field.Bool("is_spam").Optional().Nillable(),
		field.Bool("in_violation").Optional().Nillable(),
		field.Uint("violation_id").Optional().Nillable(),
		field.Bool("email_notify").Optional().Nillable(),
		field.Bool("has_media").Optional().Nillable(),
		field.String("blurhash").Optional().Nillable(),
		field.String("message").Optional().Nillable(),
		field.String("violation_header").Optional().Nillable(),
		field.String("violation_body").Optional().Nillable(),
		field.String("meta").Optional().Nillable(),
		field.String("appeal_message").Optional().Nillable(),
		field.Time("appeal_requested_at").Optional().Nillable(),
		field.Time("appeal_handled_at").Optional().Nillable(),
		field.Time("read_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Uint("severity_index").Optional().Nillable(),
		field.Uint64("thread_id").Optional().Nillable().Unique(),
		field.Time("emailed_at").Optional().Nillable(),
	}
}

// Edges of the AccountInterstitial.
func (AccountInterstitial) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("interstitials").Field("user_id").Unique(),
		edge.To("status", Status.Type).Field("item_id").Unique(),
	}
}
