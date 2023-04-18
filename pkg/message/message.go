package message

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/event"
	"github.com/go-nostr/go-nostr/pkg/filter"
)

const (
	// MessageTypeEvent represents an Event message type
	MessageTypeEvent = "EVENT"
	// MessageTypeEndOfStoredEvents represents an End of Stored Events message type
	MessageTypeEndOfStoredEvents = "EOSE"
	// MessageTypeClose represents a Close message type
	MessageTypeClose = "CLOSE"
	// MessageTypeNotice represents a Notice message type
	MessageTypeNotice = "NOTICE"
	// MessageTypeRequest represents a Request message type
	MessageTypeRequest = "REQ"
)

// NewCloseMessage creates a new CloseMessage
func NewCloseMessage(sid []byte) *CloseMessage {
	return &CloseMessage{sid}
}

// CloseMessage represents a Close message
type CloseMessage struct {
	SubscriptionID []byte
}

// Encode encodes the CloseMessage into a slice of byte slices
func (m *CloseMessage) Encode() [][]byte {
	return [][]byte{
		[]byte(MessageTypeClose),
		m.SubscriptionID,
	}
}

// NewEndOfStoredEventsMessage creates a new EndOfStoredEventsMessage
func NewEndOfStoredEventsMessage(sid []byte) *EndOfStoredEventsMessage {
	return &EndOfStoredEventsMessage{sid}
}

// EndOfStoredEventsMessage represents an End of Stored Events message
type EndOfStoredEventsMessage struct {
	SubscriptionID []byte
}

// Encode encodes the EndOfStoredEventsMessage into a slice of byte slices
func (m *EndOfStoredEventsMessage) Encode() [][]byte {
	return [][]byte{
		[]byte(MessageTypeEndOfStoredEvents),
		m.SubscriptionID,
	}
}

// NewEventMessage creates a new EventMessage
func NewEventMessage(sid []byte, e *event.Event) *EventMessage {
	return &EventMessage{
		SubscriptionID: sid,
		Event:          e,
	}
}

// EventMessage represents an Event message
type EventMessage struct {
	SubscriptionID []byte
	Event          *event.Event
}

// Encode encodes the EventMessage into a slice of byte slices
func (m *EventMessage) Encode() [][]byte {
	e, err := json.Marshal(m.Event)
	if err != nil {
		panic(err)
	}

	return [][]byte{
		[]byte(MessageTypeEvent),
		m.SubscriptionID,
		e,
	}
}

// NewNoticeMessage creates a new NoticeMessage
func NewNoticeMessage(m []byte) *NoticeMessage {
	return &NoticeMessage{m}
}

// NoticeMessage represents a Notice message
type NoticeMessage struct {
	Message []byte
}

// Encode encodes the NoticeMessage into a slice of byte slices
func (m *NoticeMessage) Encode() [][]byte {
	return [][]byte{
		[]byte(MessageTypeNotice),
		m.Message,
	}
}

// NewRequestMessage creates a new RequestMessage
func NewRequestMessage(sid []byte, f *filter.Filter) *RequestMessage {
	return &RequestMessage{
		SubscriptionID: sid,
		Filter:         f,
	}
}

// RequestMessage represents a Request message
type RequestMessage struct {
	SubscriptionID []byte
	Filter         *filter.Filter
}

// Encode encodes the RequestMessage into a slice of byte slices
func (m *RequestMessage) Encode() [][]byte {
	f, err := json.Marshal(m.Filter)
	if err != nil {
		panic(err)
	}

	return [][]byte{
		[]byte(MessageTypeRequest),
		m.SubscriptionID,
		f,
	}
}
