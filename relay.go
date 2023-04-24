package nostr

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

	"nhooyr.io/websocket"
)

const (
	defaultHostname = "0.0.0.0"
	defaultPort     = 4317
)

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

func NewRelay() *Relay {
	getInternetIdentifierHandler := &getInternetIdentifierHandler{}
	websocketHandler := &websocketHandler{}
	serveMux := &http.ServeMux{}

	serveMux.Handle("/.well-known/nostr.json", getInternetIdentifierHandler)
	serveMux.Handle("/", websocketHandler)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%+v:%+v", defaultHostname, defaultPort),
		Handler: serveMux,
	}

	return &Relay{
		conns:  make(map[*websocket.Conn]struct{}),
		server: httpServer,
	}
}

type Relay struct {
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	PubKey        string       `json:"pub_key,omitempty"`
	Contact       string       `json:"contact,omitempty"`
	SupportedNIPs []string     `json:"supported_nips,omitempty"`
	Software      string       `json:"software,omitempty"`
	Version       string       `json:"version,omitempty"`
	Limitations   *Limitations `json:"limitations,omitempty"`

	server *http.Server
	conns  map[*websocket.Conn]struct{}
	mess   chan []byte
	mu     sync.Mutex
}

func (r *Relay) ListenAndServe() error {
	return r.server.ListenAndServe()
}

func (r *Relay) Serve(l net.Listener) error {
	return r.server.Serve(l)
}

func (r *Relay) Publish(mess Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	byt, err := mess.Marshal()
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	for conn := range r.conns {
		conn.Write(ctx, websocket.MessageText, byt)
	}

	return nil
}

func (r *Relay) Subscribe(u string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		return err
	}

	r.addConn(conn)

	go r.listenConn(conn)

	return nil
}

func (r *Relay) addConn(cl *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.conns[cl] = struct{}{}
}

func (r *Relay) listenConn(conn *websocket.Conn) {
	defer r.removeConn(conn)
	for {
		_, mess, err := conn.Read(context.Background())
		if err != nil {
			fmt.Printf("Error reading from connection: %v\n", err)
			return
		}
		r.mess <- mess
	}
}

func (r *Relay) removeConn(conn *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "closing connection")
}
