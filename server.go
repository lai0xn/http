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

type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

type ResponseWriter interface {
	WriteString(str string)
	WriteStatus(status int)
	WriteHeader(key string, value string)
	WriteJson(s interface{}) error
}

// The request type
type Request struct {
	Method  string
	URL     string
	Proto   string
	Headers map[string]string
	Body    []byte
}

// The response type
type Response struct {
	Proto   string
	Status  int
	Headers map[string]string
	Body    string
}

func (w *Response) WriteString(str string) {
	w.Body = str
}

func (w *Response) WriteJson(s interface{}) error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	w.Body = string(bytes)
	return nil
}

func (w *Response) WriteHeader(key string, value string) {
	w.Headers[key] = value
}

func (w *Response) WriteStatus(status int) {
	w.Status = status
}

// httpcon makes it easier to access the handlers
type httpconn struct {
	handler Handler
	conn    net.Conn
}

func (h *httpconn) serve() {
	defer h.conn.Close()

	reader := bufio.NewReader(h.conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Error(err)
			break
		}

		parts := strings.Split(line, " ")

		req := Request{
			Method:  parts[0],
			URL:     parts[1],
			Proto:   strings.TrimSpace(parts[2]),
			Headers: make(map[string]string),
		}
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
		// reading the Body
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
		resp := Response{
			Proto:   req.Proto,
			Status:  200,
			Headers: map[string]string{"Content-Type": "text/plain"},
		}
		h.handler.ServeHTTP(&resp, &req)
		resp.Headers["Content-Length"] = strconv.Itoa(len(resp.Body))
		response := fmt.Sprintf("%s %d OK\r\n", resp.Proto, resp.Status)

		for key, value := range resp.Headers {
			response += fmt.Sprintf("%s: %s\r\n", key, value)
		}
		response += "\r\n" + resp.Body
		fmt.Println(response)
		h.conn.Write([]byte(response))
		connection := req.Headers["Connection"]
		if connection != "keep-alive" {
			break
		}
	}
}

type Server struct {
	Handler Handler
	ADDR    string
}

func (s *Server) Listen(l net.Listener) error {
	log.Info("Server Running")
	defer l.Close()
	for {

		conn, err := l.Accept()
		if err != nil {
			return err
		}
		log.Info("New Connection")
		hc := httpconn{
			handler: s.Handler,
			conn:    conn,
		}
		go hc.serve()
	}
}
