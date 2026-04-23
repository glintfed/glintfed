package data

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	App     AppConfig     `mapstructure:"app"`
	Server  ServerConfig  `mapstructure:"server"`
	Service ServiceConfig `mapstructure:"service"`
}

type AppConfig struct {
	Name        string   `mapstructure:"name" env:"APP_NAME" default:"glintfed"`
	Version     string   `mapstructure:"version" env:"APP_VERSION" default:"0.0.0"`
	Env         string   `mapstructure:"env" env:"APP_ENV"`
	Key         []byte   `mapstructure:"-"`
	KeyValue    string   `mapstructure:"key" env:"APP_KEY"`
	URL         *url.URL `mapstructure:"-"`
	URLValue    string   `mapstructure:"url" env:"APP_URL"`
	Description string   `mapstructure:"description" env:"PF_DESCRIPTION"`

	MediaTypes          string `mapstructure:"media_types" env:"MEDIA_TYPES"`
	MaxPhotoSize        int    `mapstructure:"max_photo_size" env:"MAX_PHOTO_SIZE"`
	MaxCaptionLength    int    `mapstructure:"max_caption_length" env:"MAX_CAPTION_LENGTH"`
	MaxAltextLength     int    `mapstructure:"max_altext_length" env:"PF_MEDIA_MAX_ALTTEXT_LENGTH"`
	MaxAlbumLength      int    `mapstructure:"max_album_length" env:"MAX_ALBUM_LENGTH"`
	ImageQuality        int    `mapstructure:"image_quality" env:"IMAGE_QUALITY"`
	MaxCollectionLength int    `mapstructure:"max_collection_length" env:"PF_MAX_COLLECTION_LENGTH"`
	OptimizeImage       bool   `mapstructure:"optimize_image" env:"PF_OPTIMIZE_IMAGES"`
	OptimizeVideo       bool   `mapstructure:"optimize_video" env:"PF_OPTIMIZE_VIDEOS"`
	EnforceAcountLimit  bool   `mapstructure:"enforce_account_limit" env:"LIMIT_ACCOUNT_SIZE"`
	CloudStorage        bool   `mapstructure:"cloud_storage" env:"PF_ENABLE_CLOUD"`

	MaxAvatarSize     int `mapstructure:"max_avatar_size" env:"MAX_AVATAR_SIZE"`
	MaxBioLength      int `mapstructure:"max_bio_length" env:"MAX_BIO_LENGTH"`
	MaxNameLength     int `mapstructure:"max_name_length" env:"MAX_NAME_LENGTH"`
	MinPasswordLength int `mapstructure:"min_password_length" env:"MIN_PASSWORD_LENGTH"`
	MaxAccountSize    int `mapstructure:"max_account_size" env:"MAX_ACCOUNT_SIZE"`

	Auth       AuthConfig       `mapstructure:"auth"`
	Instance   InstanceConfig   `mapstructure:"instance"`
	Federation FederationConfig `mapstructure:"federation"`
	Import     ImportConfig     `mapstructure:"import"`
	Media      MediaConfig      `mapstructure:"media"`
	Groups     GroupsConfig     `mapstructure:"groups"`
}

type AuthConfig struct {
	EnableRegistration bool        `mapstructure:"enable_registration" env:"OPEN_REGISTRATION"`
	EnableOAuth        bool        `mapstructure:"enable_oauth" env:"OAUTH_ENABLED"`
	InAppRegistration  bool        `mapstructure:"in_app_registration" env:"APP_REGISTER"`
	LoginURL           *url.URL    `mapstructure:"-"`
	LoginURLValue      string      `mapstructure:"login_url" env:"OAUTH_LOGIN_URL"`
	OAuth              OAuthConfig `mapstructure:"oauth"`
}

// OAuthConfig holds configuration for the embedded OAuth2 server.
// AccessTokenLifespan and RefreshTokenLifespan accept Go duration strings (e.g. "8760h" for 365 days).
type OAuthConfig struct {
	HMACSecret           string        `mapstructure:"hmac_secret"           env:"OAUTH_HMAC_SECRET"`
	PersonalClientID     string        `mapstructure:"personal_client_id"    env:"OAUTH_PERSONAL_CLIENT_ID"`
	AccessTokenLifespan  time.Duration `mapstructure:"access_token_lifespan" env:"OAUTH_TOKEN_EXPIRATION"`
	RefreshTokenLifespan time.Duration `mapstructure:"refresh_token_lifespan" env:"OAUTH_REFRESH_EXPIRATION"`
}

type UploaderConfig struct {
}

type ActivitypubConfig struct {
	Enabled      bool `mapstructure:"enabled" env:"ACTIVITY_PUB"`
	RemoteFollow bool `mapstructure:"remote_follow" env:"AP_REMOTE_FOLLOW"`
	SharedInbox  bool `mapstructure:"shared_inbox" env:"AP_SHAREDINBOX"`
	Inbox        bool `mapstructure:"inbox" env:"AP_INBOX"`
}

type InstanceConfig struct {
	HasLegalNotice bool           `mapstructure:"has_legal_notice" env:"INSTANCE_LEGAL_NOTICE"`
	Username       UsernameConfig `mapstructure:"username"`
	Stories        StoriesConfig  `mapstructure:"stories"`
	Label          LabelConfig    `mapstructure:"label"`
}

type UsernameConfig struct {
	Remote RemoteConfig `mapstructure:"remote"`
}

type StoriesConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

type RemoteConfig struct {
	Formats []string `mapstructure:"formats"`
	Format  string   `mapstructure:"format"`
	Custom  string   `mapstructure:"custom" env:"USERNAME_REMOTE_CUSTOM_TEXT"`
}

type ImportConfig struct {
	Instagram InstagramConfig `mapstructure:"instagram"`
}

type InstagramConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

type MediaConfig struct {
	HLS HLSConfig `mapstructure:"hls"`
}

type HLSConfig struct {
	Enabled  bool   `mapstructure:"enabled" env:"MEDIA_HLS_ENABLED"`
	Debug    bool   `mapstructure:"debug" env:"MEDIA_HLS_DEBUG"`
	P2P      bool   `mapstructure:"p2p" env:"MEDIA_HLS_P2P"`
	P2PDebug bool   `mapstructure:"p2p_debug" env:"MEDIA_HLS_P2P_DEBUG"`
	Tracker  string `mapstructure:"tracker" env:"MEDIA_HLS_P2P_TRACKER"`
	Ice      string `mapstructure:"ice" env:"MEDIA_HLS_P2P_ICE_SERVER"`
}

type LabelConfig struct {
	Covid LabelContentConfig `mapstructure:"covid"`
}

type LabelContentConfig struct {
	Enabled  bool     `mapstructure:"enabled"`
	Org      string   `mapstructure:"org"`
	URL      *url.URL `mapstructure:"-"`
	URLValue string   `mapstructure:"url"`
}

type GroupsConfig struct {
	Enabled bool `mapstructure:"enabled" env:"GROUPS_ENABLED"`
}

type ServerConfig struct {
	API APIServerConfig `mapstructure:"api"`
}

type APIServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type ServiceConfig struct {
	Database      DatabaseConfig      `mapstructure:"database"`
	OpenTelemetry OpenTelemetryConfig `mapstructure:"open_telemetry"`
}

type DatabaseConfig struct {
	SQL   SQLDBConfig `mapstructure:"sql"`
	Redis RedisConfig `mapstructure:"redis"`
}

type SQLDBConfig struct {
	Driver string `mapstructure:"driver" env:"DB_DRIVER"`
	DSN    string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr" env:"REDIS_HOST"`
	Password string `mapstructure:"password" env:"REDIS_PASSWORD"`
}

type OpenTelemetryConfig struct {
	TracingEnabled  bool   `mapstructure:"tracing_enabled"`
	TracingEndpoint string `mapstructure:"tracing_endpoint"`
}

type FederationConfig struct {
	NodeInfo        NodeInfoConfig    `mapstructure:"nodeinfo"`
	Webfinger       WebfingerConfig   `mapstructure:"webfinger"`
	NetworkTimeline bool              `mapstructure:"network_timeline" env:"PF_NETWORK_TIMELINE"`
	Activitypub     ActivitypubConfig `mapstructure:"activitypub"`
}

type NodeInfoConfig struct {
	Enabled bool `mapstructure:"enabled" env:"NODEINFO"`
}

type WebfingerConfig struct {
	Enabled bool `mapstructure:"enabled" env:"WEBFINGER"`
}

func NewConfig(name, version string, paths ...string) (*Config, error) {
	config.WithOptions(config.ParseEnv, config.ParseTime, config.ParseDefault)
	config.AddDriver(yaml.Driver)

	cfg := &Config{}

	var cfgPaths []string
	for _, p := range paths {
		stat, err := os.Stat(p)
		if err != nil {
			return nil, err
		}

		if !stat.IsDir() {
			cfgPaths = append(cfgPaths, p)
			continue
		}

		filepath.WalkDir(p, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			ext := filepath.Ext(path)
			if !d.IsDir() && (ext == ".yaml" || ext == ".yml") {
				cfgPaths = append(cfgPaths, path)
			}

			return nil
		})
	}

	if err := config.LoadFiles(cfgPaths...); err != nil {
		return nil, err
	}

	if err := config.BindStruct("", cfg); err != nil {
		return nil, err
	}

	if name != "" {
		cfg.App.Name = name
	}
	if version != "" {
		cfg.App.Version = version
	}

	if key, err := parseKey(cfg.App.KeyValue); err != nil {
		return nil, err
	} else {
		cfg.App.Key = key
	}

	if u, err := url.Parse(cfg.App.URLValue); err != nil {
		return nil, err
	} else {
		cfg.App.URL = u
	}

	if u, err := url.Parse(cfg.App.Auth.LoginURLValue); err != nil {
		return nil, err
	} else {
		cfg.App.Auth.LoginURL = u
	}

	if u, err := url.Parse(cfg.App.Instance.Label.Covid.URLValue); err != nil {
		cfg.App.Instance.Label.Covid.URL = u
	} else {
		cfg.App.Instance.Label.Covid.URL = u
	}

	return cfg, nil
}

func parseKey(value string) (ret []byte, err error) {
	if len(value) == 0 {
		return nil, errors.New("missing app.key")
	}

	switch {
	case strings.HasPrefix(value, "hex:"):
		trimmed := strings.TrimPrefix(value, "hex:")
		return hex.DecodeString(trimmed)
	case strings.HasPrefix(value, "base64:"):
		trimmed := strings.TrimPrefix(value, "base64:")
		return base64.StdEncoding.DecodeString(trimmed)
	}

	return []byte(value), nil
}
