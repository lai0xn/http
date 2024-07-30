package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (h HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	h(w, r)
}

func NewServerMux() *mux {
	return &mux{}
}

type Route struct {
	Method  string
	URL     string
	Handler Handler
}

func (m *mux) GET(route string, f HandlerFunc) {
	r := Route{
		Method:  "GET",
		URL:     route,
		Handler: Handler(f),
	}
	m.routes = append(m.routes, r)
}

func (m *mux) POST(route string, f HandlerFunc) {
	r := Route{
		Method:  "POST",
		URL:     route,
		Handler: f,
	}
	m.routes = append(m.routes, r)
}

type mux struct {
	routes []Route
}

func (m *mux) ServeHTTP(w ResponseWriter, r *Request) {
	found := false
	var h Handler

	for _, route := range m.routes {
		if route.URL == r.URL && r.Method == route.Method {
			h = route.Handler
			found = true
		}
	}
	if found {
		h.ServeHTTP(w, r)
	} else {
		w.WriteStatus(404)
	}
}
