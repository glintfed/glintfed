package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// Annotations of the Instance.
func (Instance) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "instances"},
	}
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("domain").Unique(),
		field.Bool("active_deliver").Optional().Nillable(),
		field.String("url").Optional().Nillable(),
		field.String("name").Optional().Nillable(),
		field.String("admin_url").Optional().Nillable(),
		field.String("limit_reason").Optional().Nillable(),
		field.Bool("unlisted"),
		field.Bool("auto_cw"),
		field.Bool("banned"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.String("software").Optional().Nillable(),
		field.Uint("user_count").Optional().Nillable(),
		field.Uint("status_count").Optional().Nillable(),
		field.Time("last_crawled_at").Optional().Nillable(),
		field.Time("actors_last_synced_at").Optional().Nillable(),
		field.String("notes").Optional().Nillable(),
		field.Bool("manually_added"),
		field.String("base_domain").Optional().Nillable(),
		field.Bool("ban_subdomains").Optional().Nillable(),
		field.String("ip_address").Optional().Nillable(),
		field.Bool("list_limitation"),
		field.Bool("valid_nodeinfo").Optional().Nillable(),
		field.Time("nodeinfo_last_fetched").Optional().Nillable(),
		field.Bool("delivery_timeout"),
		field.Time("delivery_next_after").Optional().Nillable(),
		field.String("shared_inbox").Optional().Nillable(),
	}
}

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profiles", Profile.Type),
	}
}
