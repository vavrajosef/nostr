package nostrgrpc

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

const (
	defaultHostname = "0.0.0.0"
	defaultPort     = 4317
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	// NOTE: TBD
	serveMux := &http.ServeMux{}

	// NOTE: TBD
	server := &http.Server{
		Handler:      serveMux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return &Server{
		server: server,
	}
}

// Serve TBD
func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", defaultHostname, defaultPort))
	if err != nil {
		panic(err)
	}

	return s.server.Serve(listener)
}
