package relay

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-nostr/nostr"
	"nhooyr.io/websocket"
)

func NewServer() *Server {
	opt := &ServerOptions{}
	serveMux := &http.ServeMux{}

	if opt.server != nil {
		opt.server.Handler = serveMux
	} else {
		opt.server = &http.Server{
			Handler: serveMux,
		}
	}

	return &Server{
		opt,
		opt.server,
	}
}

type ServerOptions struct {
	Name          []byte             `json:"name,omitempty"`
	Description   []byte             `json:"description,omitempty"`
	PubKey        []byte             `json:"pub_key,omitempty"`
	Contact       []byte             `json:"contact,omitempty"`
	SupportedNIPs []byte             `json:"supported_nips,omitempty"`
	Software      []byte             `json:"software,omitempty"`
	Version       []byte             `json:"version,omitempty"`
	Limitations   *nostr.Limitations `json:"limitations,omitempty"`

	server *http.Server
}

type Server struct {
	*ServerOptions

	server *http.Server
}

func (r *Server) Dial() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	log.Printf("Dialing...")

	c, _, err := websocket.Dial(ctx, "wss://nostr.wine", &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		panic(err)
	}

	// c.SetReadLimit(2e6)

	// defer c.Close(websocket.StatusInternalError, "the sky is falling")

	reqMessage := nostr.NewRequestMessage("npub145u3sqsd5h9:asrxasdf85q03swfs", &nostr.Filter{
		Kinds: []nostr.EventKind{nostr.EventKindShortTextNote},
		Limit: 1,
		Since: time.Now().Unix() - 1000,
	})
	byt, _ := reqMessage.Marshal()

	err = c.Write(ctx, websocket.MessageText, byt)
	if err != nil {
		panic(err)
	}

	// TODO: experiments with event un/marshaling
	// var v any
	// if err := wsjson.Read(ctx, c, &v); err != nil {
	// 	log.Printf("%+v", err)
	// }
	// log.Printf("%+v", v)

	// var eoseMessage nostr.EOSEMessage
	// _, data, _ := c.Read(ctx)
	// eoseMessage.Unmarshal(data)
	// a, _ := eoseMessage.Marshal()

	var eventMessage nostr.EventMessage
	_, data, _ := c.Read(ctx)
	eventMessage.Unmarshal(data)
	a, _ := eventMessage.Marshal()

	log.Printf("%+v", string(a))
}

func (r *Server) Serve() error {
	return r.server.ListenAndServe()
}
