package function

import (
	"context"

	"github.com/go-web/redis/connection"
)

type IFunction interface {
	SetValue(ctx context.Context, key string, value interface{}) error
	GetValue(ctx context.Context, key string) (interface{}, error)
}

func NewRedisFunction() IFunction {
	return &impl{}
}

type impl struct{}

func (*impl) SetValue(ctx context.Context, key string, value interface{}) error {
	return connection.GetRedisClient().Set(ctx, key, "value111", 0).Err()
}

func (*impl) GetValue(ctx context.Context, key string) (interface{}, error) {
	return connection.GetRedisClient().Get(ctx, key).Bytes()
}
