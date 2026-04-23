package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	entschema "entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []entschema.Annotation {
	return []entschema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("profile_id").Optional().Nillable().Unique(),
		field.String("name").Optional().Nillable(),
		field.String("username").Optional().Nillable().Unique(),
		field.String("email").Sensitive().Unique(),
		field.String("status").Optional().Nillable(),
		field.String("language").Optional().Nillable(),
		field.String("password").Sensitive(),
		field.String("remember_token").Sensitive().Optional().Nillable(),
		field.Bool("is_admin"),
		field.Time("email_verified_at").Optional().Nillable(),
		field.Time("created_at").Optional().Nillable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.Time("last_active_at").Optional().Nillable(),
		field.Bool("two_fa_enabled").StorageKey("2fa_enabled"),
		field.String("two_fa_secret").StorageKey("2fa_secret").Sensitive().Optional().Nillable(),
		field.String("two_fa_backup_codes").StorageKey("2fa_backup_codes").Sensitive().Optional().Nillable(),
		field.Time("two_fa_setup_at").StorageKey("2fa_setup_at").Optional().Nillable(),
		field.Time("delete_after").Optional().Nillable(),
		field.Bool("has_interstitial"),
		field.String("guid").Optional().Nillable().Unique(),
		field.String("domain").Optional().Nillable(),
		field.String("register_source").Optional().Nillable(),
		field.String("app_register_token").Optional().Nillable(),
		field.String("app_register_ip").Optional().Nillable(),
		field.Bool("has_roles"),
		field.Uint("parent_id").Optional().Nillable(),
		field.Uint("role_id").Optional().Nillable(),
		field.String("expo_token").Optional().Nillable(),
		field.Bool("notify_like"),
		field.Bool("notify_follow"),
		field.Bool("notify_mention"),
		field.Bool("notify_comment"),
		field.Uint64("storage_used"),
		field.Time("storage_used_updated_at").Optional().Nillable(),
		field.Bool("notify_enabled"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("filters", UserFilter.Type),
		edge.To("devices", UserDevice.Type),
		edge.To("accountLog", AccountLog.Type),
		edge.To("interstitials", AccountInterstitial.Type),
	}
}
