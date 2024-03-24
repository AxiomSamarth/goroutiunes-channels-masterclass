package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

// NewServer creates and returns a new instance of Server object.
// It also registers the handlers to the http.
func NewServer(ctx context.Context, host, port string) *Server {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	return &Server{
		Addr: fmt.Sprintf("%s:%s", host, port),
		ctx:  ctx,
	}
}

// Start starts the server and listens at the configured address.
func (s *Server) Start() error {
	log.Println("starting the server at", s.Addr)
	s.server = &http.Server{
		Addr: s.Addr,
	}
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

// Stop stops the HTTP Server gracefully.
func (s *Server) Stop() error {
	return s.server.Shutdown(s.ctx)
}
