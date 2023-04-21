package eosemessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new NoticeMessage.
func New(sid string) *NoticeMessage {
	return &NoticeMessage{sid}
}

// NoticeMessage represents an End of Stored Events message,
// signaling that there are no more events for a given subscription.
type NoticeMessage struct {
	SubscriptionID string `json:"subscription_id,omitempty"` // The ID of the subscription with no more events
}

// Marshall TBD
func (m *NoticeMessage) Marshal() ([]byte, error) {
	return json.Marshal([]string{
		string(message.TypeNotice),
		m.SubscriptionID,
	})
}
