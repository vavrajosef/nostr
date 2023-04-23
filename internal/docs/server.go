package docs

//go:generate hugo
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
	defaultPort     = 9110
)

//go:embed public
var FS embed.FS

// Server TBD
type Server struct {
	server *http.Server
}

// New TBD
func NewServer() *Server {
	// NOTE: TBD
	fs, _ := fs.Sub(FS, "public")

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

// ListenAndServe TBD
func (r *Server) ListenAndServe() error {
	return r.server.ListenAndServe()
}

// Serve TBD
func (r *Server) Serve(l net.Listener) error {
	return r.server.Serve(l)
}
