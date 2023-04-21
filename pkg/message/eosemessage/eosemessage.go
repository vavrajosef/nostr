package eosemessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new EOSEMessage.
func New(sid string) *EOSEMessage {
	return &EOSEMessage{sid}
}

// EOSEMessage represents an End of Stored Events message,
// signaling that there are no more events for a given subscription.
type EOSEMessage struct {
	SubscriptionID string `json:"subscription_id,omitempty"` // The ID of the subscription with no more events
}

// Marshall TBD
func (m *EOSEMessage) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(message.TypeEOSE),
		m.SubscriptionID,
	})
}
