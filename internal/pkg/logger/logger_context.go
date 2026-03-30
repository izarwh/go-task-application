package logger

import (
	"context"
	"log/slog"
)

type ctxKey string

const (
	RequestIDKey ctxKey = "request_id"
	UserIDKey    ctxKey = "user_id"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		r.AddAttrs(slog.String("request_id", reqID))
	}
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		r.AddAttrs(slog.String("user_id", userID))
	}
	return h.Handler.Handle(ctx, r)
}
