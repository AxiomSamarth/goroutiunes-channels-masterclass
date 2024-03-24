package server

import (
	"context"
	"net/http"
)

// Server is the wrapper for the upstream http.Server
// to consume with implementation convenience in the context
// of this project.
type Server struct {
	// Addr represents the address at which the server will serve and listen.
	Addr string

	// ctx is the context passed to and referred by every individual instance of this server.
	ctx context.Context

	// server is the native HTTP Server object that offers & manages the server operations.
	server *http.Server
}
