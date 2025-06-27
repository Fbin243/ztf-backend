package auth

import (
	"context"
	"errors"
)

var AuthKey = "userId"

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
