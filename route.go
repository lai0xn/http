package http

import "net/http"

// Route defines a route with an HTTP method, URL pattern, and associated handler.
type Route struct {
	method  string
	url     string
	handler Handler
}

func (r *Route) match(w http.ResponseWriter, req *http.Request) bool {
	if req.URL.String() != r.url {
		return false
	}
	if req.Method != r.method && r.method != "ALL" {
		return false
	}
	return true

}

func (r *Route) GetHandler() Handler {
	return r.handler
}

func (r *Route) GetURL() string {
	return r.url
}
