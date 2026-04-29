package cache

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisDrv struct {
	client *redis.Client
}

func NewRedisDriver(client *redis.Client) Driver {
	return &redisDrv{
		client: client,
	}
}

func (d *redisDrv) Has(ctx context.Context, key string) bool {
	n, err := d.client.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return n > 0
}

func (d *redisDrv) Get(ctx context.Context, key string) any {
	val, err := d.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return nil
	}
	if decoded, ok := decodeRedisValue(val); ok {
		return decoded
	}
	return val
}

func (d *redisDrv) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	encoded, err := encodeRedisValue(val)
	if err != nil {
		return err
	}
	return d.client.Set(ctx, key, encoded, ttl).Err()
}

func (d *redisDrv) Del(ctx context.Context, key string) error {
	return d.client.Del(ctx, key).Err()
}

func (d *redisDrv) Clear(ctx context.Context) error {
	return d.client.FlushDB(ctx).Err()
}

func encodeRedisValue(val any) (any, error) {
	switch val.(type) {
	case string, []byte:
		return val, nil
	}

	if val != nil {
		gob.Register(val)
	}
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(&val); err != nil {
		return nil, err
	}
	return buf.String(), nil
}

func decodeRedisValue(text string) (any, bool) {
	var val any
	if err := gob.NewDecoder(bytes.NewBufferString(text)).Decode(&val); err != nil {
		return nil, false
	}
	return val, true
}
