package client_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/internal/client"
	"nhooyr.io/websocket"
)

// TODO: mock websocket conn

func TestNewClient(t *testing.T) {
	conn, _, err := websocket.Dial(context.TODO(), "wss://nostr.wine", &websocket.DialOptions{})
	if err != nil {
		t.Error(err)
	}
	type args struct {
		conn *websocket.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SHOULD create client",
			args: args{
				conn: conn,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := client.NewClient(tt.args.conn)

			if cl == nil {
				t.Fatalf("expected %+v, to be not nil", cl)
			}

			t.Log(cl)
		})
	}
}

func TestClient_PublishMessage(t *testing.T) {
	conn, _, err := websocket.Dial(context.TODO(), "wss://nostr.wine", &websocket.DialOptions{})
	if err != nil {
		t.Error(err)
	}
	type fields struct {
		conn *websocket.Conn
	}
	type args struct {
		mess nostr.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "SHOULD publish request message",
			fields: fields{
				conn: conn,
			},
			args: args{
				mess: nostr.NewRequestMessage("asdf", &nostr.Filter{}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := client.NewClient(tt.fields.conn)
			err := cl.PublishMessage(tt.args.mess)

			if err != nil {
				t.Fatal(err)
			}

			_, byt, err := conn.Read(context.TODO())
			if err != nil {
				t.Fatal(err)
			}

			var data any
			err = json.Unmarshal(byt, &data)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%+v", data)

			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
