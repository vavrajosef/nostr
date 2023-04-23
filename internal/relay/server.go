package relay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-nostr/nostr"
	"nhooyr.io/websocket"
)

const (
	defaultHostname = "0.0.0.0"
	defaultPort     = 4317
)

type getHealthHandler struct {
}

func (h *getHealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(struct {
		Status    string `json:"status"`
		Timestamp int64  `json:"timestamp"`
	}{
		Status:    "OK",
		Timestamp: time.Now().Unix(),
	})

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

type getInternetIdentifierHandler struct {
}

func (h *getInternetIdentifierHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(struct {
		Names  []string `json:"names,omitempty"`
		Relays []string `json:"relays,omitempty"`
	}{
		Names:  []string{"bob"},
		Relays: []string{},
	})

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

type websocketHandler struct {
}

func (h *websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	// TODO: add service call to add subscriber

	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		log.Printf("%v", err)
		return
	}
}

func NewServer() *Server {
	// NOTE: handlers
	getHealthHandler := &getHealthHandler{}
	getInternetIdentifierHandler := &getInternetIdentifierHandler{}
	websocketHandler := &websocketHandler{}

	// NOTE: mux
	serveMux := &http.ServeMux{}

	// NOTE: routes
	serveMux.Handle("/health", getHealthHandler)
	serveMux.Handle("/.well-known/nostr.json", getInternetIdentifierHandler)
	serveMux.Handle("/nostr", websocketHandler)

	// NOTE: http server
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%+v:%+v", defaultHostname, defaultPort),
		Handler: serveMux,
	}

	return &Server{
		server: httpServer,
	}
}

type Server struct {
	Name          string             `json:"name,omitempty"`
	Description   string             `json:"description,omitempty"`
	PubKey        string             `json:"pub_key,omitempty"`
	Contact       string             `json:"contact,omitempty"`
	SupportedNIPs []string           `json:"supported_nips,omitempty"`
	Software      string             `json:"software,omitempty"`
	Version       string             `json:"version,omitempty"`
	Limitations   *nostr.Limitations `json:"limitations,omitempty"`

	server *http.Server
}

func (r *Server) ListenAndServe() error {
	return r.server.ListenAndServe()
}

func (r *Server) Serve(l net.Listener) error {
	return r.server.Serve(l)
}
