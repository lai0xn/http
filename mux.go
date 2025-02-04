package http

// Router is a simple multiplexer for routing HTTP requests to handlers.
// It maintains a list of routes and middlewares to handle incoming requests.
type Router struct {
	routes      []*Route         // List of registered routes
	middlewares []MiddlewareFunc // List of middlewares to be applied to handlers
}

// MiddlewareFunc is a type definition for middleware functions.
// Middleware functions wrap around handlers to provide additional functionality.
type MiddlewareFunc func(next Handler) Handler

// NewServerMux creates and returns a new instance of Router.
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

// Use adds a middleware function to the list of middlewares.
// Middlewares are applied in the order they are added.
func (m *Router) Use(middleware MiddlewareFunc) {
	m.middlewares = append(m.middlewares, middleware)
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

// applyMiddleware applies all registered middlewares to the given handler.
// Middlewares are applied in the order they were added.
func (m *Router) applyMiddleware(handler Handler) Handler {
	for _, m := range m.middlewares {
		handler = m(handler)
	}
	return handler
}

// ServeHTTP matches the request URL and method against the registered routes.
// If a match is found, it applies middlewares and invokes the corresponding handler.
// If no route matches, it returns a 404 Not Found status.
func (m *Router) ServeHTTP(w ResponseWriter, r *Request) {
	for _, route := range m.routes {
		if route.match(r) {
			handler := m.applyMiddleware(route.handler)
			handler.ServeHTTP(w, r)
			return
		}
	}
	// Return 404 Not Found if no route matches the request.
	w.WriteStatus(404)
}
