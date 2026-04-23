package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// WebsocketsStatisticsEntry holds the schema definition for the WebsocketsStatisticsEntry entity.
type WebsocketsStatisticsEntry struct {
	ent.Schema
}

// Annotations of the WebsocketsStatisticsEntry.
func (WebsocketsStatisticsEntry) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "websockets_statistics_entries"},
	}
}

// Fields of the WebsocketsStatisticsEntry.
func (WebsocketsStatisticsEntry) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("id"),
		field.String("app_id"),
		field.Int("peak_connection_count"),
		field.Int("websocket_message_count"),
		field.Int("api_message_count"),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
	}
}

// Edges of the WebsocketsStatisticsEntry.
func (WebsocketsStatisticsEntry) Edges() []ent.Edge {
	return []ent.Edge{}
}
