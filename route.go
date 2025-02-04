package http

// Route defines a route with an HTTP method, URL pattern, and associated handler.
type Route struct {
	method  string  // HTTP method (e.g., GET, POST, PUT, DELETE, PATCH, or "ALL" for any method)
	url     string  // URL pattern to match against incoming requests
	handler Handler // Handler function to execute when the route is matched
}

// match checks if the incoming request matches the route's method and URL.
// It returns true if the request matches, otherwise false.
func (r *Route) match(req *Request) bool {
	// Check if the request URL matches the route's URL.
	if req.URL != r.url {
		return false
	}
	// Check if the request method matches the route's method.
	// If the route's method is "ALL", it matches any HTTP method.
	if req.Method != r.method && r.method != "ALL" {
		return false
	}
	return true
}

// GetHandler returns the handler associated with the route.
func (r *Route) GetHandler() Handler {
	return r.handler
}

// GetURL returns the URL pattern associated with the route.
func (r *Route) GetURL() string {
	return r.url
}
