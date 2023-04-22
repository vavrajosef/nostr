package client

import (
	"context"
	"time"

	"github.com/go-nostr/nostr"
	"nhooyr.io/websocket"
)

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

type Client struct {
	conn *websocket.Conn
}

func (cl *Client) PublishMessage(mess nostr.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	byt, err := mess.Marshal()
	if err != nil {
		return err
	}

	return cl.conn.Write(ctx, websocket.MessageText, byt)
}
