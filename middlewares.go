package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func RequestLoggerMiddleware(l zerolog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		// Logs all responses.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := NewWrapResponseWriter(w, r.ProtoMajor)
			rec := httptest.NewRecorder()
			ctx := r.Context()
			path := r.URL.EscapedPath()
			reqData, _ := httputil.DumpRequest(r, false)
			logger := l.Log().Timestamp().Str("path", path).Bytes("request_data", reqData)
			defer func(begin time.Time) {
				status := ww.Status()
				tookMs := time.Since(begin).Milliseconds()
				logger.Int64("took", tookMs).Int("status_code", status).Msgf("[%d] %s http request for %s took %dms", status, r.Method, path, tookMs)
			}(time.Now())

			// Replace "logger" with a custom type, like ContextKey("logger")
			ctx = context.WithValue(ctx, "logger", logger)
			next.ServeHTTP(rec, r.WithContext(ctx))
			// this copies the recorded response to the response writer
			for k, v := range rec.Header() {
				ww.Header()[k] = v
			}
			ww.WriteHeader(rec.Code)
			rec.Body.WriteTo(ww)
		})
	}
}
