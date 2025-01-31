package http

// Router is a simple multiplexer for routing HTTP requests to handlers.
type Router struct {
	routes []*Route
}

func NewServerMux() *Router {
	return &Router{}
}

// GET registers a new GET route with the specified URL and handler function.
func (m *Router) GET(route string, f HandlerFunc) {
	r := &Route{
		method:  "GET",
		url:     route,
		handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// POST registers a new POST route with the specified URL and handler function.
func (m *Router) POST(route string, f HandlerFunc) {
	r := &Route{
		method:  "POST",
		url:     route,
		handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// PUT registers a new PUT route with the specified URL and handler function.
func (m *Router) PUT(route string, f HandlerFunc) {
	r := &Route{
		method:  "PUT",
		url:     route,
		handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// DELETE registers a new DELETE route with the specified URL and handler function.
func (m *Router) DELETE(route string, f HandlerFunc) {
	r := &Route{
		method:  "DELETE",
		url:     route,
		handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// PATCH registers a new PATCH route with the specified URL and handler function.
func (m *Router) PATCH(route string, f HandlerFunc) {
	r := &Route{
		method:  "PATCH",
		url:     route,
		handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// ServeHTTP matches the request URL and method against the registered routes
// and invokes the corresponding handler if a match is found.
func (m *Router) ServeHTTP(w ResponseWriter, r *Request) {
	for _, route := range m.routes {
		if route.match(r) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// Return 404 Not Found if no route matches the request.
	w.WriteStatus(404)
}
