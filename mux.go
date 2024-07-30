package http

// HandlerFunc is a function type that implements the Handler interface.
type HandlerFunc func(w ResponseWriter, r *Request)

// ServeHTTP calls the HandlerFunc with the given ResponseWriter and Request.
func (h HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	h(w, r)
}

// NewServerMux creates and returns a new instance of mux.
func NewServerMux() *mux {
	return &mux{}
}

// Route defines a route with an HTTP method, URL pattern, and associated handler.
type Route struct {
	Method  string
	URL     string
	Handler Handler
}

// GET registers a new GET route with the specified URL and handler function.
func (m *mux) GET(route string, f HandlerFunc) {
	r := Route{
		Method:  "GET",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// POST registers a new POST route with the specified URL and handler function.
func (m *mux) POST(route string, f HandlerFunc) {
	r := Route{
		Method:  "POST",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// PUT registers a new PUT route with the specified URL and handler function.
func (m *mux) PUT(route string, f HandlerFunc) {
	r := Route{
		Method:  "PUT",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// DELETE registers a new DELETE route with the specified URL and handler function.
func (m *mux) DELETE(route string, f HandlerFunc) {
	r := Route{
		Method:  "DELETE",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// PATCH registers a new PATCH route with the specified URL and handler function.
func (m *mux) PATCH(route string, f HandlerFunc) {
	r := Route{
		Method:  "PATCH",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

// mux is a simple multiplexer for routing HTTP requests to handlers.
type mux struct {
	routes []Route
}

// ServeHTTP matches the request URL and method against the registered routes
// and invokes the corresponding handler if a match is found.
func (m *mux) ServeHTTP(w ResponseWriter, r *Request) {
	for _, route := range m.routes {
		if route.URL == r.URL && r.Method == route.Method {
			route.Handler.ServeHTTP(w, r)
			return
		}
	}
	// Return 404 Not Found if no route matches the request.
	w.WriteStatus(404)
}
