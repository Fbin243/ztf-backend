package auth

import (
	"context"
	"errors"
)

type ContextKey string

var AuthKey ContextKey = "userId"

func GetAuthKey(ctx context.Context) (int64, error) {
	userId, exists := ctx.Value(AuthKey).(int64)
	if !exists || userId == 0 {
		return 0, errors.New("auth key not found in context")
	}

	return userId, nil
}

func SetAuthKey(ctx context.Context, userId int64) context.Context {
	if userId == 0 {
		return ctx
	}

	return context.WithValue(ctx, AuthKey, userId)
}
