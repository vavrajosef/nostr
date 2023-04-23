package nostr_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-nostr/nostr"
	"nhooyr.io/websocket"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "SHOULD create client",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := nostr.NewClient()

			if cl == nil {
				t.Fatalf("expected %+v, to be not nil", cl)
			}

			t.Log(cl)
		})
	}
}

func TestClient_Publish(t *testing.T) {
	type fields struct {
		u string
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
				u: "wss://relay.example.com", // Replace with a valid relay URL
			},
			args: args{
				mess: nostr.NewRequestMessage("asdf", &nostr.Filter{}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := nostr.NewClient()

			go cl.Subscribe(tt.fields.u)

			err := cl.Publish(tt.args.mess)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestClient_Subscribe(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "SHOULD subscribe to relay",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				c, err := websocket.Accept(w, r, nil)
				if err != nil {
					t.Fatal(err)
				}
				defer c.Close(websocket.StatusNormalClosure, "")

				for {
					_, byt, err := c.Read(r.Context())
					if err != nil {
						return
					}
					msg := &nostr.RequestMessage{}
					err = msg.Unmarshal(byt)
					if err != nil {
						return
					}
					// TODO: complete implementation
					// if msg.SubscriptionID != "asdf-1234" {
					// 	return
					// }
					// if reflect.DeepEqual(msg.Filter, &nostr.Filter{}) {
					// 	t.Fatalf("expected filter: %+v, got filter: %+v", msg.Filter, &nostr.Filter{})
					// 	return
					// }
					fmt.Printf("Received message: %v\n", msg)
				}
			}))
			defer server.Close()

			u := "ws" + server.URL[4:]

			cl := nostr.NewClient()
			err := cl.Subscribe(u)
			if err != nil {
				t.Fatal(err)
			}

			time.Sleep(1 * time.Second)
			err = cl.Publish(nostr.NewRequestMessage("asdf-1234", &nostr.Filter{}))
			if err != nil {
				t.Fatal(err)
			}

			time.Sleep(1 * time.Second)
		})
	}
}
