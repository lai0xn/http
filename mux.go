package http

// Router is a simple multiplexer for routing HTTP requests to handlers.
type Router struct {
	routes []*Route
}

func NewServerMux() *Router {
	return &Router{}
}

// Route defines a route with an HTTP method, URL pattern, and associated handler.
type Route struct {
	Method  string
	URL     string
	Handler Handler
}

// GET registers a new GET route with the specified URL and handler function.
func (m *Router) GET(route string, f HandlerFunc) {
	r := &Route{
		Method:  "GET",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// POST registers a new POST route with the specified URL and handler function.
func (m *Router) POST(route string, f HandlerFunc) {
	r := &Route{
		Method:  "POST",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// PUT registers a new PUT route with the specified URL and handler function.
func (m *Router) PUT(route string, f HandlerFunc) {
	r := &Route{
		Method:  "PUT",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// DELETE registers a new DELETE route with the specified URL and handler function.
func (m *Router) DELETE(route string, f HandlerFunc) {
	r := &Route{
		Method:  "DELETE",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// PATCH registers a new PATCH route with the specified URL and handler function.
func (m *Router) PATCH(route string, f HandlerFunc) {
	r := &Route{
		Method:  "PATCH",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// ServeHTTP matches the request URL and method against the registered routes
// and invokes the corresponding handler if a match is found.
func (m *Router) ServeHTTP(w ResponseWriter, r *Request) {
	for _, route := range m.routes {
		if route.URL == r.URL && r.Method == route.Method {
			route.Handler.ServeHTTP(w, r)
			return
		}
	}
	// Return 404 Not Found if no route matches the request.
	w.WriteStatus(404)
}
