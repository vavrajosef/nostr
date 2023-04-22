package relay

import (
	"net/http"

	"github.com/go-nostr/nostr"
)

func NewServer() *Server {
	opt := &ServerOptions{}
	serveMux := &http.ServeMux{}

	if opt.server != nil {
		opt.server.Handler = serveMux
	} else {
		opt.server = &http.Server{
			Handler: serveMux,
		}
	}

	return &Server{
		opt,
		opt.server,
	}
}

type ServerOptions struct {
	Name          []byte             `json:"name,omitempty"`
	Description   []byte             `json:"description,omitempty"`
	PubKey        []byte             `json:"pub_key,omitempty"`
	Contact       []byte             `json:"contact,omitempty"`
	SupportedNIPs []byte             `json:"supported_nips,omitempty"`
	Software      []byte             `json:"software,omitempty"`
	Version       []byte             `json:"version,omitempty"`
	Limitations   *nostr.Limitations `json:"limitations,omitempty"`

	server *http.Server
}

type Server struct {
	*ServerOptions

	server *http.Server
}

func (r *Server) Serve() error {
	return r.server.ListenAndServe()
}
