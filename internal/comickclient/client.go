package comickclient

import (
	"context"
	"net/http"
	"time"
)

type Client struct {
	HTTP *http.Client
}

func New(opts ...func(*http.Client)) *Client {
	c := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   30 * time.Second,
	}

	for _, o := range opts {
		o(c)
	}

	return &Client{HTTP: c}
}


func (c *Client) Download(ctx context.Context, id, path string) error {
	return nil
}