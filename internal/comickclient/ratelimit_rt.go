package comickclient

import (
	"net/http"
	"strings"

	"golang.org/x/time/rate"
)

type ratelimitRT struct {
	transport http.RoundTripper
	limits    map[string]*rate.Limiter
}

func (rl ratelimitRT) RoundTrip(r *http.Request) (*http.Response, error) {
	str := r.URL.Host + r.URL.Path
	for part, lim := range rl.limits {
		if strings.Contains(str, part) {
			err := lim.Wait(r.Context())
			if err != nil {
				return nil, err
			}
			break
		}
	}
	return rl.transport.RoundTrip(r)
}

func WithRateLimits(l map[string]*rate.Limiter) func(*http.Client) {
	return func(c *http.Client) {
		c.Transport = ratelimitRT{
			transport: c.Transport,
			limits:    l,
		}
	}
}
