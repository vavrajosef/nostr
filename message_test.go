package nostr_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/go-nostr/nostr"
)

func TestMessageEvent_Marshal(t *testing.T) {
	tests := []struct {
		name     string
		message  *nostr.EventMessage
		expected string
	}{
		{
			name: "TestMessageEvent_Marshal_Valid",
			message: &nostr.EventMessage{
				Type:           nostr.MessageTypeEvent,
				SubscriptionID: "test-subscription-id",
				Event: &nostr.MetadataEvent{
					Kind: nostr.EventKindMetadata,
				},
			},
			expected: `["EVENT","test-subscription-id",{"kind":0}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.message.Marshal()
			if err != nil {
				t.Errorf("nostr.EventMessage.Marshal() error = %v", err)
			}

			if string(result) != tt.expected {
				t.Errorf("nostr.EventMessage.Marshal() = %v, expected %v", string(result), tt.expected)
			}
		})
	}
}

func TestMessageEvent_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedEvent *nostr.EventMessage
	}{
		{
			name:  "TestMessageEvent_UnmarshalJSON_Valid",
			input: `["EVENT","test-subscription-id",{"kind": 0}]`,
			expectedEvent: &nostr.EventMessage{
				Type:           nostr.MessageTypeEvent,
				SubscriptionID: "test-subscription-id",
				Event: &nostr.MetadataEvent{
					Kind: nostr.EventKindMetadata,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventMessage := &nostr.EventMessage{}
			err := json.Unmarshal([]byte(tt.input), eventMessage)

			if err != nil {
				t.Errorf("nostr.EventMessage.UnmarshalJSON() error = %v", err)
			}

			if !reflect.DeepEqual(eventMessage, tt.expectedEvent) {
				t.Errorf("nostr.EventMessage.UnmarshalJSON() = %+v, expected %+v", eventMessage, tt.expectedEvent)
			}
		})
	}
}
