package countmessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new CountMessage.
func New(sid string) *CountMessage {
	return &CountMessage{sid}
}

// CountMessage TBD
type CountMessage struct {
	SubscriptionID string `json:"subscription_id,omitempty"` // The ID of the subscription to be closed
}

// Marshall TBD
func (m *CountMessage) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(message.TypeCount),
		m.SubscriptionID,
	})
}
