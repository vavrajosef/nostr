package message

// Type represents the type of a message.
type Type string

const (
	// TypeAuth represents an authentication message type.
	TypeAuth Type = "AUTH"
	// TypeClose represents a Close message type.
	TypeClose Type = "CLOSE"
	// TypeCount represents a count message type, usually for counting events.
	TypeCount Type = "COUNT"
	// TypeEndOfStoredEvents represents an End of Stored Events message type.
	TypeEOSE Type = "EOSE"
	// TypeEvent represents an Event message type.
	TypeEvent Type = "EVENT"
	// TypeNotice represents a Notice message type, usually for notifications.
	TypeNotice Type = "NOTICE"
	// TypeOk represents a success confirmation message type.
	TypeOk Type = "OK"
	// TypeRequest represents a Request message type.
	TypeRequest Type = "REQ"
)

// Message is an interface for encoding and marshaling messages.
type Message interface {
	Marshal() ([]byte, error)
}
