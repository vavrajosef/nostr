package amounttag

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/tag"
)

// New TBD
func New(eventID string, relayURL string, marker tag.Marker) *EventTag {
	return &EventTag{
		Type:     tag.TypeEvent,
		EventID:  eventID,
		RelayURL: relayURL,
		Marker:   marker,
	}
}

// EventTag is a tag for an event, including the event ID and relay URL.
type EventTag struct {
	Type     tag.Type   `json:"type,omitempty"`      // The type of the tag, which is "e" for event tags
	EventID  string     `json:"event_id,omitempty"`  // The unique identifier for the associated event
	RelayURL string     `json:"relay_url,omitempty"` // The URL of the relay server for the event
	Marker   tag.Marker `json:"marker,omitempty"`    // The URL of the relay server for the event
}

// Marshall TBD
func (t *EventTag) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(tag.TypeAmount),
		t.EventID,
		t.RelayURL,
		string(t.Marker),
	})
}
