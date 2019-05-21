package httputil

import "net/http"

// Doer represents a component able to perform an HTTP request
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
