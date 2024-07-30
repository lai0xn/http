# HTTP

## Overview

This is a simple HTTP server library (didn't find a cool name yet) written in Go. It includes functionalities for handling HTTP requests and responses, routing, and creating an HTTP server.

### Installation
```bash
go get github.com/lai0xn/http
```


### Example Usage

The `Handler` interface defines a method to serve HTTP requests.

```go
func main() {
	mux := http.NewServerMux()

	mux.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteString("Hello, World!")
	})

	mux.POST("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteJson(map[string]string{"echo": string(r.Body)})
	})

	server := http.Server{
		Handler: mux,
		ADDR:    ":8080",
	}

	listener, err := net.Listen("tcp", server.ADDR)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	err = server.Listen(listener)
	if err != nil {
		fmt.Println("Error while running server:", err)
	}
}
```

### Contribution
Feel free to contribute to this library by submitting issues or pull requests
