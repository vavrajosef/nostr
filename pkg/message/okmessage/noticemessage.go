package eosemessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new OkMessage.
func New(sid string) *OkMessage {
	return &OkMessage{sid}
}

// OkMessage represents an End of Stored Events message,
// signaling that there are no more events for a given subscription.
type OkMessage struct {
	SubscriptionID string `json:"subscription_id,omitempty"` // The ID of the subscription with no more events
}

// Marshall TBD
func (m *OkMessage) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(message.TypeOk),
		m.SubscriptionID,
	})
}
