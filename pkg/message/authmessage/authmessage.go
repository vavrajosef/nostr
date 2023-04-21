package authmessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new AuthMessage.
func New(sid string) *AuthMessage {
	return &AuthMessage{sid}
}

// AuthMessage represents a Auth message, used to close a subscription.
type AuthMessage struct {
	SubscriptionID string `json:"subscription_id,omitempty"` // The ID of the subscription to be closed
}

// Marshall TBD
func (m *AuthMessage) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(message.TypeAuth),
		m.SubscriptionID,
	})
}
