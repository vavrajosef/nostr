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
