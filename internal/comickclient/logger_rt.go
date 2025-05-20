package comickclient

import (
	"log/slog"
	"net/http"
	"time"
)

type loggerRT struct {
	transport http.RoundTripper
	logger    *slog.Logger
}

func (l loggerRT) RoundTrip(r *http.Request) (*http.Response, error) {

	start := time.Now()
	resp, err := l.transport.RoundTrip(r)
	code := 0
	if resp != nil {
		code = resp.StatusCode
	}
	level := slog.LevelDebug
	if err != nil || code > 400 {
		level = slog.LevelError
	}

	l.logger.LogAttrs(
		r.Context(), level, "http",
		slog.String("method", r.Method),
		slog.String("url", r.URL.String()),
		slog.Int("status", code),
		slog.Duration("elapsed", time.Since(start)),
		slog.Any("err", err),
	)

	return resp, err
}

func WithLogger(l *slog.Logger) func(*http.Client) {
	return func(c *http.Client) {
		c.Transport = loggerRT{
			transport: c.Transport,
			logger:    l,
		}
	}
}
