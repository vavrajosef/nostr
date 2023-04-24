package nostr

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

//go:generate npm run build -w web

//go:embed internal/web/dist
var web embed.FS

// NewClient TBD
func NewClient() *Client {
	dist, err := fs.Sub(web, "internal/web/dist")
	if err != nil {
		fmt.Printf("Error reading from fs: %+v", err)
		return nil
	}
	return &Client{
		fs:    dist,
		mess:  make(chan []byte),
		conns: make(map[*websocket.Conn]struct{}),
	}
}

// Client TBD
type Client struct {
	conns map[*websocket.Conn]struct{}
	fs    fs.FS
	mess  chan []byte
	mu    sync.Mutex
}

// TODO: add event handlers
func (cl *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(cl.fs)).ServeHTTP(w, r)
}

func (cl *Client) Publish(mess Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	byt, err := mess.Marshal()
	if err != nil {
		return err
	}

	cl.mu.Lock()
	defer cl.mu.Unlock()

	for conn := range cl.conns {
		conn.Write(ctx, websocket.MessageText, byt)
	}
	return nil
}

func (cl *Client) Subscribe(u string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		return err
	}

	cl.addConn(conn)

	go cl.listenConn(conn)

	return nil
}

func (cl *Client) addConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	cl.conns[conn] = struct{}{}
}

func (cl *Client) listenConn(conn *websocket.Conn) {
	defer cl.removeConn(conn)

	for {
		_, mess, err := conn.Read(context.Background())
		if err != nil {
			fmt.Printf("Error reading from relay: %v\n", err)
			return
		}

		cl.mess <- mess
	}
}

func (cl *Client) removeConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	delete(cl.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "closing connection")
}
