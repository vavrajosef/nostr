package message

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/event"
	"github.com/go-nostr/go-nostr/pkg/filter"
)

// Type represents the type of a message.
type Type string

const (
	// TypeAuth TBD
	TypeAuth Type = "AUTH"
	// TypeClose represents a Close message type.
	TypeClose Type = "CLOSE"
	// TypeCount TBD
	TypeCount Type = "COUNT"
	// TypeEndOfStoredEvents represents an End of Stored Events message type
	TypeEndOfStoredEvents Type = "EOSE"
	// TypeEvent represents an Event message type
	TypeEvent Type = "EVENT"
	// TypeNotice represents a Notice message type
	TypeNotice Type = "NOTICE"
	// TypeOK TBD
	TypeOk Type = "OK"
	// TypeRequest represents a Request message type
	TypeRequest Type = "REQ"
)

// Message TBD
type Message interface {
	Encode() ([][]byte, error)
}

// NewAuthMessage creates a new AuthMessage.
func NewAuthMessage(subID []byte) *AuthMessage {
	return &AuthMessage{subID}
}

// AuthMessage represents a Auth message, used to close a subscription.
type AuthMessage struct {
	SubscriptionID []byte `json:"subscription_id,omitempty"` // The ID of the subscription to be closed
}

// Encode encodes the AuthMessage into a slice of byte slices.
func (m *AuthMessage) Encode() ([][]byte, error) {
	return [][]byte{
		[]byte(TypeAuth),
		m.SubscriptionID,
	}, nil
}

// NewCountMessage creates a new CountMessage.
func NewCountMessage(subID []byte) *CountMessage {
	return &CountMessage{subID}
}

// CountMessage represents a Count message, used to close a subscription.
type CountMessage struct {
	SubscriptionID []byte `json:"subscription_id,omitempty"` // The ID of the subscription to be closed
}

// Encode encodes the CountMessage into a slice of byte slices.
func (m *CountMessage) Encode() ([][]byte, error) {
	return [][]byte{
		[]byte(TypeCount),
		m.SubscriptionID,
	}, nil
}

// NewCloseMessage creates a new CloseMessage.
func NewCloseMessage(subID []byte) *CloseMessage {
	return &CloseMessage{subID}
}

// CloseMessage represents a Close message, used to close a subscription.
type CloseMessage struct {
	SubscriptionID []byte `json:"subscription_id,omitempty"` // The ID of the subscription to be closed
}

// Encode encodes the CloseMessage into a slice of byte slices.
func (m *CloseMessage) Encode() ([][]byte, error) {
	return [][]byte{
		[]byte(TypeClose),
		m.SubscriptionID,
	}, nil
}

// NewEndOfStoredEventsMessage creates a new EndOfStoredEventsMessage.
func NewEndOfStoredEventsMessage(subID []byte) *EndOfStoredEventsMessage {
	return &EndOfStoredEventsMessage{subID}
}

// EndOfStoredEventsMessage represents an End of Stored Events message,
// signaling that there are no more events for a given subscription.
type EndOfStoredEventsMessage struct {
	SubscriptionID []byte `json:"subscription_id,omitempty"` // The ID of the subscription with no more events
}

// Encode encodes the EndOfStoredEventsMessage into a slice of byte slices.
func (m *EndOfStoredEventsMessage) Encode() ([][]byte, error) {
	return [][]byte{
		[]byte(TypeEndOfStoredEvents),
		m.SubscriptionID,
	}, nil
}

// NewEventMessage creates a new EventMessage.
func NewEventMessage(subID []byte, e *event.Event) *EventMessage {
	return &EventMessage{
		SubscriptionID: subID,
		Event:          e,
	}
}

// EventMessage represents an Event message, containing a single event.
type EventMessage struct {
	SubscriptionID []byte       `json:"subscription_id,omitempty"` // The ID of the subscription for the event
	Event          *event.Event `json:"event,omitempty"`           // The event object
}

// Encode encodes the EventMessage into a slice of byte slices.
func (m *EventMessage) Encode() ([][]byte, error) {
	e, err := json.Marshal(m.Event)
	if err != nil {
		return nil, err
	}

	return [][]byte{
		[]byte(TypeEvent),
		m.SubscriptionID,
		e,
	}, nil
}

// NewNoticeMessage creates a new NoticeMessage.
func NewNoticeMessage(m []byte) *NoticeMessage {
	return &NoticeMessage{m}
}

// NoticeMessage represents a Notice message, which provides additional information.
type NoticeMessage struct {
	Message []byte `json:"message,omitempty"` // The content of the notice message
}

// Encode encodes the NoticeMessage into a slice of byte slices.
func (m *NoticeMessage) Encode() ([][]byte, error) {
	return [][]byte{
		[]byte(TypeNotice),
		m.Message,
	}, nil
}

// NewRequestMessage creates a new RequestMessage.
func NewRequestMessage(subID []byte, f *filter.Filter) *RequestMessage {
	return &RequestMessage{
		SubscriptionID: subID,
		Filter:         f,
	}
}

// RequestMessage represents a Request message, which is used to request events
// that match a specified filter.
type RequestMessage struct {
	SubscriptionID []byte         `json:"subscription_id,omitempty"` // The ID of the subscription for the request
	Filter         *filter.Filter `json:"filter,omitempty"`          // The filter used to request events
}

// Encode encodes the RequestMessage into a slice of byte slices.
func (m *RequestMessage) Encode() ([][]byte, error) {
	f, err := json.Marshal(m.Filter)
	if err != nil {
		return nil, err
	}

	return [][]byte{
		[]byte(TypeRequest),
		m.SubscriptionID,
		f,
	}, nil
}
