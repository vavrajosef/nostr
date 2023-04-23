package client

//go:generate npm run build
import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"time"
)

const (
	defaultHostname = "0.0.0.0"
	defaultPort     = 4200
)

//go:embed dist
var FS embed.FS

// Server TBD
type Server struct {
	server *http.Server
}

// New TBD
func NewServer() *Server {
	// NOTE: TBD
	fs, _ := fs.Sub(FS, "dist")

	// NOTE: TBD
	serveMux := &http.ServeMux{}

	// NOTE: TBD
	serveMux.Handle("/", http.FileServer(http.FS(fs)))

	// NOTE: TBD
	server := &http.Server{
		Addr:         fmt.Sprintf("%+v:%+v", defaultHostname, defaultPort),
		Handler:      serveMux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return &Server{
		server: server,
	}
}

// ListenAndServe ...
func (r *Server) ListenAndServe() error {
	return r.server.ListenAndServe()
}

// Serve TBD
func (r *Server) Serve(l net.Listener) error {
	return r.server.Serve(l)
}
