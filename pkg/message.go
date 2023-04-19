package nostr

import (
	"encoding/json"
)

// Type represents the type of a message.
type MessageType string

const (
	// MessageTypeAuth TBD
	MessageTypeAuth MessageType = "AUTH"
	// MessageTypeClose represents a Close message type.
	MessageTypeClose MessageType = "CLOSE"
	// MessageTypeCount TBD
	MessageTypeCount MessageType = "COUNT"
	// MessageTypeEndOfStoredEvents represents an End of Stored Events message type
	MessageTypeEndOfStoredEvents MessageType = "EOSE"
	// MessageTypeEvent represents an Event message type
	MessageTypeEvent MessageType = "EVENT"
	// MessageTypeNotice represents a Notice message type
	MessageTypeNotice MessageType = "NOTICE"
	// MessageTypeOK TBD
	MessageTypeOk MessageType = "OK"
	// MessageTypeRequest represents a Request message type
	MessageTypeRequest MessageType = "REQ"
)

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
		[]byte(MessageTypeAuth),
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
		[]byte(MessageTypeCount),
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
		[]byte(MessageTypeClose),
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
		[]byte(MessageTypeEndOfStoredEvents),
		m.SubscriptionID,
	}, nil
}

// NewEventMessage creates a new EventMessage.
func NewEventMessage(subID []byte, e *Event) *EventMessage {
	return &EventMessage{
		SubscriptionID: subID,
		Event:          e,
	}
}

// EventMessage represents an Event message, containing a single event.
type EventMessage struct {
	SubscriptionID []byte `json:"subscription_id,omitempty"` // The ID of the subscription for the event
	Event          *Event `json:"event,omitempty"`           // The event object
}

// Encode encodes the EventMessage into a slice of byte slices.
func (m *EventMessage) Encode() ([][]byte, error) {
	e, err := json.Marshal(m.Event)
	if err != nil {
		return nil, err
	}

	return [][]byte{
		[]byte(MessageTypeEvent),
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
		[]byte(MessageTypeNotice),
		m.Message,
	}, nil
}

// NewRequestMessage creates a new RequestMessage.
func NewRequestMessage(subID []byte, f *Filter) *RequestMessage {
	return &RequestMessage{
		SubscriptionID: subID,
		Filter:         f,
	}
}

// RequestMessage represents a Request message, which is used to request events
// that match a specified filter.
type RequestMessage struct {
	SubscriptionID []byte  `json:"subscription_id,omitempty"` // The ID of the subscription for the request
	Filter         *Filter `json:"filter,omitempty"`          // The filter used to request events
}

// Encode encodes the RequestMessage into a slice of byte slices.
func (m *RequestMessage) Encode() ([][]byte, error) {
	f, err := json.Marshal(m.Filter)
	if err != nil {
		return nil, err
	}

	return [][]byte{
		[]byte(MessageTypeRequest),
		m.SubscriptionID,
		f,
	}, nil
}
