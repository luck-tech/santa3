package contexts

import (
	"context"

	"github.com/labstack/echo/v4"
)

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

const (
	RequestID ContextKey = "x-request-id"
	UserID    ContextKey = "x-user-id"
	SessionID ContextKey = "x-session-id"
)

var contextkeys = []ContextKey{
	RequestID,
}

func GetRequestID(ctx context.Context) string {
	v, ok := ctx.Value(RequestID.String()).(string)
	if !ok {
		return ""
	}
	return v
}

func GetUserID(ctx context.Context) string {
	v, ok := ctx.Value(UserID.String()).(string)
	if !ok {
		return ""
	}
	return v
}

func GetSessionID(ctx context.Context) string {
	v, ok := ctx.Value(SessionID.String()).(string)
	if !ok {
		return ""
	}
	return v
}

// ConvertContext はecho.Contextのkeyをcopyしてcontext.Contextに変換
func ConvertContext(c echo.Context) context.Context {
	ctx := context.Background()
	for _, key := range contextkeys {
		v := c.Get(key.String())
		ctx = context.WithValue(ctx, key, v)
	}

	return ctx
}
