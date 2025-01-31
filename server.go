package http

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

// Handler interface defines the method to serve HTTP requests.
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

// ResponseWriter interface defines methods for writing responses.
type ResponseWriter interface {
	WriteString(str string)
	WriteStatus(status int)
	WriteHeader(key string, value string)
	WriteJson(s interface{}) error
}

// Request represents an HTTP request.
type Request struct {
	Method  string
	URL     string
	Proto   string
	Headers map[string]string
	Body    []byte
}

// Response represents an HTTP response.
type Response struct {
	Proto   string
	Status  int
	Headers map[string]string
	Body    string
}

// WriteString sets the response body to a string value.
func (w *Response) WriteString(str string) {
	w.Body = str
}

// WriteJson serializes an object as JSON and sets it as the response body.
func (w *Response) WriteJson(s interface{}) error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	w.Body = string(bytes)
	return nil
}

// WriteHeader sets a header key-value pair in the response.
func (w *Response) WriteHeader(key string, value string) {
	w.Headers[key] = value
}

// WriteStatus sets the response status code.
func (w *Response) WriteStatus(status int) {
	w.Status = status
}

// httpconn manages the connection and handles HTTP requests.
type httpconn struct {
	handler Handler
	conn    net.Conn
}

// serve handles the HTTP requests and responses for the connection.
func (h *httpconn) serve() {
	defer h.conn.Close()

	reader := bufio.NewReader(h.conn)

	for {
		// Read the request line (method, URL, and protocol)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Error(err)
			break
		}

		parts := strings.Split(line, " ")

		// Initialize the request object
		req := Request{
			Method:  parts[0],
			URL:     parts[1],
			Proto:   strings.TrimSpace(parts[2]),
			Headers: make(map[string]string),
		}

		// Read headers
		for {
			headerLine, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			if headerLine == "\r\n" {
				break
			}
			headerParts := strings.SplitN(headerLine, ":", 2)
			if len(headerParts) == 2 {
				req.Headers[strings.TrimSpace(headerParts[0])] = strings.TrimSpace(headerParts[1])
			}
		}

		// Read the body if Content-Length is specified
		if bodyLen, ok := req.Headers["Content-Length"]; ok {
			length, err := strconv.Atoi(bodyLen)
			if err != nil {
				return
			}
			body := make([]byte, length)
			_, err = reader.Read(body)
			if err != nil {
				return
			}
			req.Body = body
		}

		// Prepare and handle the response
		resp := Response{
			Proto:   req.Proto,
			Status:  200,
			Headers: map[string]string{"Content-Type": "text/plain"},
		}
		h.handler.ServeHTTP(&resp, &req)
		resp.Headers["Content-Length"] = strconv.Itoa(len(resp.Body))

		// Construct the response
		response := fmt.Sprintf("%s %d OK\r\n", resp.Proto, resp.Status)
		for key, value := range resp.Headers {
			response += fmt.Sprintf("%s: %s\r\n", key, value)
		}
		response += "\r\n" + resp.Body

		// Send the response
		fmt.Println(response)
		h.conn.Write([]byte(response))

		// Check if connection should be kept alive
		connection := req.Headers["Connection"]
		if connection != "keep-alive" {
			break
		}
	}
}

// Server represents an HTTP server.
type Server struct {
	Handler Handler
	ADDR    string
}

// Listen starts the server and handles incoming connections.
func (s *Server) Listen(l net.Listener) error {
	log.Info("Server Running")
	defer l.Close()

	for {
		// Accept new connections
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		log.Info("New Connection")

		// Handle the connection in a new goroutine
		hc := httpconn{
			handler: s.Handler,
			conn:    conn,
		}
		go hc.serve()
	}
}

// HandlerFunc is a function type that implements the Handler interface.
type HandlerFunc func(w ResponseWriter, r *Request)

// ServeHTTP calls the HandlerFunc with the given ResponseWriter and Request.
func (h HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	h(w, r)
}
