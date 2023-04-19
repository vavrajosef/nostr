package nostr

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

func New(opt *Options) *Relay {
	return &Relay{}
}

// NOTE: copied from cmd/chat.go

// Subscriber represents a subscriber.
// Messages are sent on the msgs channel and if the client
// cannot keep up with the messages, closeSlow is called.
type Subscriber struct {
	msgs      chan []byte
	closeSlow func()
}

type Relay struct {
	*Options

	// subscriberMessageBuffer controls the max number
	// of messages that can be queued for a subscriber
	// before it is kicked.
	//
	// Defaults to 16.
	subscriberMessageBuffer int

	// publishLimiter controls the rate limit applied to the publish endpoint.
	//
	// Defaults to one publish every 100ms with a burst of 8.
	publishLimiter *rate.Limiter

	// logf controls where logs are sent.
	// Defaults to log.Printf.
	logf func(f string, v ...interface{})

	// serveMux routes the various endpoints to the appropriate handler.
	serveMux http.ServeMux

	subscribersMu sync.Mutex
	subscribers   map[*Subscriber]struct{}
}
