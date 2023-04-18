package web

import (
	"embed"
	"io/fs"
	"net"
	"net/http"
	"time"
)

//go:embed dist
var FS embed.FS

// Server TBD
type Server struct {
	server *http.Server
}

// NewServer TBD
func NewServer() *Server {
	// NOTE: TBD
	static, _ := fs.Sub(FS, "dist")

	// NOTE: TBD
	serveMux := &http.ServeMux{}

	// NOTE: TBD
	serveMux.Handle("/", http.FileServer(http.FS(static)))

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
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	return s.server.Serve(l)
}
