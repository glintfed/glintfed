package client

import (
	"glintfed/ent"
	"glintfed/ent/enttest"
	"glintfed/internal/data"
	"glintfed/pkg/cache"
	"testing"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/XSAM/otelsql"
	"github.com/go-redis/redismock/v9"
	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
)

type Database struct {
	Ent     *ent.Client
	RDB     *redis.Client
	RDBMock redismock.ClientMock
}

func NewDatabase(cfg *data.Config) (db *Database, err error) {
	db = &Database{}

	if err = db.initSQLClient(cfg); err != nil {
		return
	}

	if err = db.initRedisClient(cfg); err != nil {
		return
	}

	return
}

func NewTestDatabase(t *testing.T) (db *Database, err error) {
	db = &Database{
		Ent: enttest.Open(t,
			"sqlite3", "file:ent?mode=memory&_fk=1",
			enttest.WithOptions(ent.Log(t.Log)),
		),
	}

	db.RDB, db.RDBMock = redismock.NewClientMock()

	return
}

func (c *Database) initSQLClient(cfg *data.Config) error {
	db, err := otelsql.Open(
		cfg.Service.Database.SQL.Driver, cfg.Service.Database.SQL.DSN,
		otelsql.WithAttributes(semconv.DBSystemNameKey.String(cfg.Service.Database.SQL.Driver)),
	)
	if err != nil {
		return err
	}

	c.Ent = ent.NewClient(ent.Driver(entsql.OpenDB(cfg.Service.Database.SQL.Driver, db)))
	return db.Ping()
}

func (c *Database) initRedisClient(cfg *data.Config) (err error) {
	c.RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Service.Database.Redis.Addr,
		Password: cfg.Service.Database.Redis.Password,
	})

	cache.Register(cache.NewRedisDriver(c.RDB))

	return redisotel.InstrumentTracing(c.RDB)
}
