package eventtag

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/tag"
)

// New TBD
func New(eventID []byte, relayURL []byte) *DelegationTag {
	return &DelegationTag{
		Type:     tag.TypeEvent,
		EventID:  eventID,
		RelayURL: relayURL,
	}
}

// DelegationTag is a tag for an event, including the event ID and relay URL.
type DelegationTag struct {
	Type     tag.Type `json:"type,omitempty"`      // The type of the tag, which is "e" for event tags
	EventID  []byte   `json:"event_id,omitempty"`  // The unique identifier for the associated event
	RelayURL []byte   `json:"relay_url,omitempty"` // The URL of the relay server for the event
	Marker   []byte   `json:"marker,omitempty"`    // The URL of the relay server for the event
}

// Encode encodes the DelegationTag into a byte slice.
func (t *DelegationTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(tag.TypeEvent))
	b = append(b, t.EventID)
	b = append(b, t.RelayURL)

	return b
}

func (t *DelegationTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}
