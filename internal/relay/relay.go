package relay

import (
	"context"
	"log"
	"time"

	"github.com/go-nostr/go-nostr"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func NewRelay(opt *Options) (*Relay, error) {
	return &Relay{opt}, nil
}

type Options struct {
	Name          []byte             `json:"name,omitempty"`
	Description   []byte             `json:"description,omitempty"`
	PubKey        []byte             `json:"pub_key,omitempty"`
	Contact       []byte             `json:"contact,omitempty"`
	SupportedNIPs []byte             `json:"supported_nips,omitempty"`
	Software      []byte             `json:"software,omitempty"`
	Version       []byte             `json:"version,omitempty"`
	Limitations   *nostr.Limitations `json:"limitations,omitempty"`
}

type Relay struct {
	*Options
}

func (r *Relay) Dial() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	log.Printf("Dialing...")

	c, _, err := websocket.Dial(ctx, "wss://eden.nostr.land", &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	// byt, _ := json.Marshal()

	// log.Printf("%v", string(byt))

	// err = wsjson.Write(ctx, c, byt)
	// if err != nil {
	// 	panic(err)
	// }

	// TODO: replace with appropriate interface
	v := []any{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		panic(err)
	}

	log.Printf("%v", v)

	c.Close(websocket.StatusNormalClosure, "")
}
