package mangaeden

import "net/http"

// Doer exposes a Do function
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is a mangaeden implementation of a Doer
type Client struct {
	c Doer
}
