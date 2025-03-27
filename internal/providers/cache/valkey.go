package cache

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/valkey-io/valkey-go"
)

type Valkey interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	Flush(ctx context.Context) error
	Ping(ctx context.Context) error
}

type valkeyImpl struct {
	client valkey.Client
}

func (c *valkeyImpl) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	strValue, ok := value.(string)

	if !ok {
		jsonValue, err := json.Marshal(value)

		if err != nil {
			return errors.New("value must be a string")
		}

		strValue = string(jsonValue)
	}

	if expiration == 0 {
		expiration = time.Hour
	}

	cmd := c.client.B().Set().Key(key).Value(strValue).ExSeconds(int64(expiration.Seconds())).Build()
	return c.client.Do(ctx, cmd).Error()
}

func (c *valkeyImpl) Get(ctx context.Context, key string) (string, error) {
	cmd := c.client.B().Get().Key(key).Build()
	result, err := c.client.Do(ctx, cmd).ToString()
	if err != nil {
		if errors.Is(err, valkey.Nil) {
			return "", errors.New("key not found")
		}
		return "", err
	}
	return result, nil
}

func (c *valkeyImpl) Delete(ctx context.Context, key string) error {
	cmd := c.client.B().Del().Key(key).Build()
	return c.client.Do(ctx, cmd).Error()
}

func (c *valkeyImpl) Flush(ctx context.Context) error {
	cmd := c.client.B().Flushall().Build()
	return c.client.Do(ctx, cmd).Error()
}

func (c *valkeyImpl) Ping(ctx context.Context) error {
	cmd := c.client.B().Ping().Build()
	return c.client.Do(ctx, cmd).Error()
}

func New(address string) Valkey {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{address},
	})
	if err != nil {
		slog.Error("Error connecting to Valkey", "error", err)
		os.Exit(1)
	}
	return &valkeyImpl{client: client}
}
