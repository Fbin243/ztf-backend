package auth

import (
	"context"
	"errors"
)

type ContextKey string

var AuthKey ContextKey = "userId"

func GetAuthKey(ctx context.Context) (string, error) {
	userId, exists := ctx.Value(AuthKey).(string)
	if !exists || userId == "" {
		return "", errors.New("auth key not found in context")
	}

	return userId, nil
}

func SetAuthKey(ctx context.Context, userId string) context.Context {
	if userId == "" {
		return ctx
	}

	return context.WithValue(ctx, AuthKey, userId)
}
