package message

import (
	"github.com/go-nostr/go-nostr/event"
	"github.com/go-nostr/go-nostr/filter"
)

const (
	MessageTypeEvent   = "EVENT"
	MessageTypeRequest = "REQ"
	MessageTypeClose   = "CLOSE"
)

type Message interface {
	Encode() []byte
}

type CloseMessage struct {
	SubscriptionID string
}

type EventMessage struct {
	SubscriptionID string
	Event          *event.Event
}

type RequestMessage struct {
	SubscriptionID string         `json:"subscriptionId"`
	Filter         *filter.Filter `json:"filter"`
}
