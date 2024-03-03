package ctxutil

import (
	"context"

	"github.com/taniko/blog/internal/domain/model/user/vo"
)

type key[T any] struct{}

func WithValue[T any](ctx context.Context, val T) context.Context {
	return context.WithValue(ctx, key[T]{}, val)
}

func Value[T any](ctx context.Context) (T, bool) {
	val, ok := ctx.Value(key[T]{}).(T)
	return val, ok
}

// GetUserID ユーザーIDを取得する
func GetUserID(ctx context.Context) (vo.UserID, bool) {
	return Value[vo.UserID](ctx)
}
