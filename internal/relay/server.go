package relay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
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

	w.Header().Add("Access-Control-Allow-Origin", "*")
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

	w.Header().Add("Access-Control-Allow-Origin", "*")
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
		clients: make(map[*client]struct{}),
		server:  httpServer,
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

	server  *http.Server
	clients map[*client]struct{}
	mess    chan []byte
	mu      sync.Mutex
}

type client struct {
	conn *websocket.Conn
}

func (r *Server) ListenAndServe() error {
	return r.server.ListenAndServe()
}

func (r *Server) Serve(l net.Listener) error {
	return r.server.Serve(l)
}

func (r *Server) Publish(mess nostr.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	byt, err := mess.Marshal()
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	for cl := range r.clients {
		cl.conn.Write(ctx, websocket.MessageText, byt)
	}

	return nil
}

func (r *Server) Subscribe(u string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		return err
	}

	cl := &client{conn: conn}
	r.addClient(cl)

	go r.listenClient(cl)

	return nil
}

func (r *Server) addClient(cl *client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clients[cl] = struct{}{}
}

func (r *Server) listenClient(cl *client) {
	defer r.removeClient(cl)

	for {
		_, mess, err := cl.conn.Read(context.Background())
		if err != nil {
			fmt.Printf("Error reading from client: %v\n", err)
			return
		}

		r.mess <- mess
	}
}

func (r *Server) removeClient(cl *client) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.clients, cl)
	cl.conn.Close(websocket.StatusNormalClosure, "closing connection")
}
