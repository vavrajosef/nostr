package eosemessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new EventMessage.
func New(sid string) *EventMessage {
	return &EventMessage{sid}
}

// EventMessage represents an End of Stored Events message,
// signaling that there are no more events for a given subscription.
type EventMessage struct {
	SubscriptionID string `json:"subscription_id,omitempty"` // The ID of the subscription with no more events
}

// Marshall TBD
func (m *EventMessage) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(message.TypeEvent),
		m.SubscriptionID,
	})
}
