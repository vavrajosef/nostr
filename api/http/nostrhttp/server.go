package nostrhttp

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-nostr/go-nostr/internal/core"
)

const (
	defaultHostname = "0.0.0.0"
	defaultPort     = 4318
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	// NOTE: TBD
	serveMux := &http.ServeMux{}

	svc := core.NewService()

	getHealthHandler := newGetHealthHandler(svc)
	getInternetIdentifier := newGetInternetIdentifier(svc)

	serveMux.Handle("/health", getHealthHandler)
	serveMux.Handle("/.well-known/nostr.json", getInternetIdentifier)

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
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", defaultHostname, defaultPort))
	if err != nil {
		panic(err)
	}

	return s.server.Serve(lis)
}
