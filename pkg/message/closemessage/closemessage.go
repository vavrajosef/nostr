package closemessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new CloseMessage.
func New(sid string) *CloseMessage {
	return &CloseMessage{sid}
}

// CloseMessage represents a Close message, used to close a subscription.
type CloseMessage struct {
	SubscriptionID string `json:"subscription_id,omitempty"` // The ID of the subscription to be closed
}

// Marshall TBD
func (m *CloseMessage) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(message.TypeClose),
		m.SubscriptionID,
	})
}
