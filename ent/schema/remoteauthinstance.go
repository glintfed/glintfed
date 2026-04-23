package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RemoteAuthInstance holds the schema definition for the RemoteAuthInstance entity.
type RemoteAuthInstance struct {
	ent.Schema
}

// Annotations of the RemoteAuthInstance.
func (RemoteAuthInstance) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "remote_auth_instances"},
	}
}

// Fields of the RemoteAuthInstance.
func (RemoteAuthInstance) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("domain").Optional().Nillable().Unique(),
		field.Uint("instance_id").Optional().Nillable(),
		field.String("client_id").Optional().Nillable(),
		field.String("client_secret").Optional().Nillable(),
		field.String("redirect_uri").Optional().Nillable(),
		field.String("root_domain").Optional().Nillable(),
		field.Bool("allowed").Optional().Nillable(),
		field.Bool("banned"),
		field.Bool("active"),
		field.Time("last_refreshed_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the RemoteAuthInstance.
func (RemoteAuthInstance) Edges() []ent.Edge {
	return []ent.Edge{}
}
