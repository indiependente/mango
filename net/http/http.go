package http

import "net/http"

// Doer represents a component able to do an HTTP request
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is an HTTP client
type Client struct {
	cl http.Client
}

// NewClient returns a Client pointer
func NewClient(hc http.Client) *Client {
	return &Client{
		cl: hc,
	}
}

// Do executes an HTTP request
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.cl.Do(req)
}
