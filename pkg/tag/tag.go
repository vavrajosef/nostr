package tag

// Type represents the type of a tag.
type Type string

// Tag is an interface for different types of tags.
type Tag interface {
	Encode() []byte
}

const (
	// TypeEvent represents an event tag.
	TypeEvent Type = "e"
	// TypePubKey represents a public key tag.
	TypePubKey Type = "p"
)

// EventTag is a tag for an event, including the event ID and relay URL.
type EventTag struct {
	Type     string `json:"type,omitempty"`      // The type of the tag, which is "e" for event tags
	EventID  []byte `json:"event_id,omitempty"`  // The unique identifier for the associated event
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the event
}

// Encode encodes the EventTag into a byte slice.
func (t *EventTag) Encode() []byte {
	data := make([]byte, 0)
	data = append(data, []byte(TypeEvent)...)
	data = append(data, t.EventID...)
	data = append(data, t.RelayURL...)

	return data
}

// PubKeyTag is a tag for a public key, including the public key and relay URL.
type PubKeyTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the PubKeyTag into a byte slice.
func (t *PubKeyTag) Encode() []byte {
	data := make([]byte, 0)
	data = append(data, []byte(TypePubKey)...)
	data = append(data, t.PubKey...)
	data = append(data, t.RelayURL...)

	return data
}
