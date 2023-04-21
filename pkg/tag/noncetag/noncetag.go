package noncetag

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/tag"
)

// New TBD
func New(eventID []byte, relayURL []byte) *NonceTag {
	return &NonceTag{
		Type:     tag.TypeNonce,
		NonceID:  eventID,
		RelayURL: relayURL,
	}
}

// NonceTag is a tag for an event, including the event ID and relay URL.
type NonceTag struct {
	Type     tag.Type `json:"type,omitempty"`      // The type of the tag, which is "e" for event tags
	NonceID  []byte   `json:"event_id,omitempty"`  // The unique identifier for the associated event
	RelayURL []byte   `json:"relay_url,omitempty"` // The URL of the relay server for the event
	Marker   []byte   `json:"marker,omitempty"`    // The URL of the relay server for the event
}

// Encode encodes the NonceTag into a byte slice.
func (t *NonceTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(tag.TypeNonce))
	b = append(b, t.NonceID)
	b = append(b, t.RelayURL)

	return b
}

func (t *NonceTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}
