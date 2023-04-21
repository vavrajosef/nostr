package requestmessage

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/filter"
	"github.com/go-nostr/go-nostr/pkg/message"
)

// New creates a new RequestMessage.
func New(sid string, f *filter.Filter) *RequestMessage {
	return &RequestMessage{
		SubscriptionID: sid,
		Filter:         f,
	}
}

// RequestMessage represents a Request message, which is used to request events
// that match a specified filter.
type RequestMessage struct {
	SubscriptionID string         `json:"subscription_id,omitempty"` // The ID of the subscription for the request
	Filter         *filter.Filter `json:"filter,omitempty"`          // The filter used to request events
}

// Marshal TBD
func (m *RequestMessage) Marshal() ([]byte, error) {
	f, err := json.Marshal(m.Filter)
	if err != nil {
		return nil, err
	}

	return json.Marshal([]string{
		string(message.TypeRequest),
		m.SubscriptionID,
		string(f),
	})
}
