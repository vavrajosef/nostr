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

func NewCloseMessage(opt *CloseMessageOptions) *CloseMessage {
	return &CloseMessage{}
}

type CloseMessageOptions struct {
	SubscriptionID string
}

type CloseMessage struct{}

func NewEventMessage(opt *EventMessageOptions) *EventMessage {
	return &EventMessage{}
}

type EventMessageOptions struct {
	SubscriptionID string
	Event          *event.Event
}

type EventMessage struct{}

func NewRequestMessage(opt *RequestMessageOptions) *RequestMessage {
	return &RequestMessage{}
}

type RequestMessageOptions struct {
	SubscriptionID string
	Filter         *filter.Filter
}

type RequestMessage struct{}
