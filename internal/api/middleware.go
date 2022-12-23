package api

import (
	"context"
	"net/http"

	"github.com/polygonid/sh-id-platform/internal/log"
)

func LogMiddleware(ctx context.Context) StrictMiddlewareFunc {
	return func(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
		return func(ctxReq context.Context, w http.ResponseWriter, r *http.Request, args interface{}) (interface{}, error) {
			return f(log.CopyFromContext(ctx, ctxReq), w, r, args)
		}
	}
}
