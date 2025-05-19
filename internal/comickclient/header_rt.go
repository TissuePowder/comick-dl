package comickclient

import (
	"net/http"
	"slices"
)

type headerRT struct {
	transport http.RoundTripper
	header http.Header
}

func (h headerRT) RoundTrip(r *http.Request) (*http.Response, error) {
	rc := r.Clone(r.Context())
	rc.Header = rc.Header.Clone()

	for k, v := range h.header {
		rc.Header[k] = slices.Clone(v)
	}
	return h.transport.RoundTrip(rc)
}


func WithHeaders(h http.Header) func(*http.Client) {
	hc := make(http.Header, len(h))
	for k, v := range h {
		hc[k] = slices.Clone(v)
	}

	return func(c *http.Client) {
		c.Transport = &headerRT{
			transport: c.Transport,
			header: hc,
		}
	}

}