package main

import (
	"log"
	"net"

	"github.com/lai0xn/http-server"
)

func main() {
	// with handler

	// server := http.Server{
	// Handler: h{},
	// }
	m := http.NewServerMux()

	m.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteString("Hello World \n")
	})
	m.GET("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteStatus(http.StatusBadRequest)
	})
	server := http.Server{
		Handler: m,
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	server.Listen(l)
}

// type h struct{}

// example handler usage

//func (s h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	if r.URL == "/name" && r.Method == "POST" {
//		var b Body
//		json.Unmarshal(r.Body, &b)
//		w.WriteHeader("Content-Type", "application/json")
//		w.WriteStatus(201)
//		w.WriteString("your name is " + b.Name + "\n")
//		return
//	}
//	w.WriteString("There is nothing here\n")
//}
