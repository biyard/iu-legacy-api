package redis

import (
	"context"
	"encoding/json"

	"biyard.co/common/base"
	"biyard.co/common/logger"
)

type Document[T any] struct {
	cli *Redis
}

func NewDocument[T any](cli *Redis) *Document[T] {
	return &Document[T]{cli: cli}
}

func (r *Document[T]) Cache(ctx context.Context, key string, value T) error {
	data, err := json.Marshal(value)
	if err != nil {
		logger.Errorf(ctx, "error marshalling document: %v", err)
		return base.ErrCacheMarshal
	}

	if err := r.cli.Cache(ctx, key, string(data)); err != nil {
		logger.Errorf(ctx, "error caching document: %v", err)
		return base.ErrCachePush
	}

	return nil
}

func (r *Document[T]) Delete(ctx context.Context, key string) error {
	if err := r.cli.Delete(ctx, key); err != nil {
		logger.Errorf(ctx, "error deleting document: %v", err)
		return base.ErrCacheDelete
	}

	return nil
}

func (r *Document[T]) Get(ctx context.Context, key string) (*T, error) {
	val, err := r.cli.Get(ctx, key)
	if err != nil {
		logger.Errorf(ctx, "error getting document: %v", err)
		return nil, base.ErrCacheGet
	}

	var doc T

	if err := json.Unmarshal([]byte(val), &doc); err != nil {
		logger.Errorf(ctx, "error unmarshalling document: %v", err)
		return nil, base.ErrCacheUnmarshal
	}

	return &doc, nil
}
