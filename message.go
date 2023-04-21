package nostr

import (
	"encoding/json"
	"errors"
)

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

// EventMessage represents a message containing an event.
type EventMessage struct {
	Type  MessageType `json:"type,omitempty"`
	Event Event       `json:"event,omitempty"`
}

// Marshal marshals the EventMessage into a JSON byte array.
func (m *EventMessage) Marshal() ([]byte, error) {
	arr := make([]interface{}, 2)

	arr[0] = string(m.Type)
	arr[1] = m.Event

	return json.Marshal(arr)
}

// UnmarshalJSON unmarshals a JSON byte array into an EventMessage.
func (m *EventMessage) UnmarshalJSON(data []byte) error {
	arr := make([]json.RawMessage, 3)

	if err := json.Unmarshal(data, &arr); err != nil {
		return err
	}

	if err := json.Unmarshal(arr[0], &m.Type); err != nil {
		return err
	}

	var eventMap map[string]interface{}
	if err := json.Unmarshal(arr[2], &eventMap); err != nil {
		return err
	}

	// Determine the event type based on the eventMap and unmarshal it accordingly
	if eventType, ok := eventMap["kind"].(float64); ok {
		eventKind := EventKind(uint32(eventType))
		var eventInstance Event

		switch eventKind {
		case EventKindMetadata:
			eventInstance = &MetadataEvent{}
		default:
			return errors.New("unsupported event kind")
		}

		if err := json.Unmarshal(arr[2], eventInstance); err != nil {
			return err
		}

		m.Event = eventInstance
	} else {
		return errors.New("event kind is not a number")
	}

	return nil
}
