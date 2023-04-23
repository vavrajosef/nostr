package nostr

import (
	"context"
	"fmt"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

func NewClient() *Client {
	return &Client{
		msg:    make(chan []byte),
		relays: make(map[*relay]struct{}),
	}
}

type Client struct {
	msg    chan []byte
	mu     sync.Mutex
	relays map[*relay]struct{}
}

type relay struct {
	conn *websocket.Conn
}

// TODO: add event handlers

func (cl *Client) Publish(mess Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	byt, err := mess.Marshal()
	if err != nil {
		return err
	}

	cl.mu.Lock()
	defer cl.mu.Unlock()
	for r := range cl.relays {
		r.conn.Write(ctx, websocket.MessageText, byt)
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

	r := &relay{conn: conn}
	cl.addRelay(r)

	go cl.listenRelay(r)

	return nil
}

func (cl *Client) addRelay(r *relay) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	cl.relays[r] = struct{}{}
}

func (cl *Client) listenRelay(r *relay) {
	defer cl.removeRelay(r)

	for {
		_, msg, err := r.conn.Read(context.Background())
		if err != nil {
			fmt.Printf("Error reading from relay: %v\n", err)
			return
		}

		cl.msg <- msg
	}
}

func (cl *Client) removeRelay(r *relay) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	delete(cl.relays, r)
	r.conn.Close(websocket.StatusNormalClosure, "closing connection")
}
