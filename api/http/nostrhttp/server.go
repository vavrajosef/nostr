package nostrhttp

import (
	"encoding/json"
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

func newGetHealthHandler(svc *core.Service) *getHealthHandler {
	return &getHealthHandler{svc}
}

type getHealthHandler struct {
	svc *core.Service
}

func (h *getHealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &core.GetHealthRequest{}
	res, _ := h.svc.GetHealth(req)
	byt, _ := json.Marshal(res)

	// NOTE: TBD
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	// NOTE: TBD
	w.Write(byt)
}

func newGetInternetIdentifier(svc *core.Service) *getInternetIdentifier {
	return &getInternetIdentifier{svc}
}

type getInternetIdentifier struct {
	svc *core.Service
}

func (h *getInternetIdentifier) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, _ := h.svc.GetInternetIdentifier(&core.GetInternetIdentifierRequest{})
	byt, _ := json.Marshal(res)

	// NOTE: TBD
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	// NOTE: TBD
	w.Write(byt)
}

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
