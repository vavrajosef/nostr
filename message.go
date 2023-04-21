package nostr

import "encoding/json"

// MessageType represents the type of a message.
type MessageType string

const (
	// MessageTypeAuth represents an authentication message type.
	MessageTypeAuth MessageType = "AUTH"
	// MessageTypeClose represents a Close message type.
	MessageTypeClose MessageType = "CLOSE"
	// MessageTypeCount represents a count message type, usually for counting events.
	MessageTypeCount MessageType = "COUNT"
	// MessageTypeEndOfStoredEvents represents an End of Stored Events message type.
	MessageTypeEOSE MessageType = "EOSE"
	// MessageTypeEvent represents an Event message type.
	MessageTypeEvent MessageType = "EVENT"
	// MessageTypeNotice represents a Notice message type, usually for notifications.
	MessageTypeNotice MessageType = "NOTICE"
	// MessageTypeOk represents a success confirmation message type.
	MessageTypeOk MessageType = "OK"
	// MessageTypeRequest represents a Request message type.
	MessageTypeRequest MessageType = "REQ"
)

// Message is an interface for encoding and marshaling messages.
type Message interface {
	Marshal() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type AuthMessage struct{}

func (m *AuthMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *AuthMessage) UnmarshalJSON(data []byte) error {
	json.Unmarshal(data, m)

	return nil
}
