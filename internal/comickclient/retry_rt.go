package comickclient

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type retryRT struct {
	transport  http.RoundTripper
	delay      time.Duration
	maxRetries int
}

func (rt retryRT) RoundTrip(r *http.Request) (*http.Response, error) {

	var resp *http.Response
	var err error

	for attempt := range rt.maxRetries {
		ctxErr := r.Context().Err()
		if ctxErr != nil {
			return nil, ctxErr
		}

		resp, err = rt.transport.RoundTrip(r)

		if err == nil {
			code := resp.StatusCode

			if code < 400 {
				return resp, nil
			}

			if code < 500 {
				// err = fmt.Errorf("http %d %s", code, http.StatusText(code))
				slog.Error(
					"client error (no retry)",
					"method", r.Method,
					"url", r.URL.String(),
					"status", code,
					"error", err,
				)
				resp.Body.Close()
				return nil, err
			}
		}

		if attempt == rt.maxRetries-1 {
			break
		}

		time.Sleep(rt.delay)
	}

	slog.Error(
		"retry: giving up",
		"method", r.Method,
		"url", r.URL.String(),
		"tries", rt.maxRetries,
		"last_error", err,
	)

	os.Exit(1)

	// for compilation. the code below is never reached
	return nil, context.Canceled

}

func WithRetry(d time.Duration, m int) func(*http.Client) {
	return func(c *http.Client) {
		c.Transport = retryRT{
			transport:  c.Transport,
			delay:      d,
			maxRetries: m,
		}
	}
}
