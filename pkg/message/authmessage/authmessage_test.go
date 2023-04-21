package authmessage_test

import (
	"bytes"
	"testing"

	"github.com/go-nostr/go-nostr/pkg/message/authmessage"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			subID string
		}
		expected *authmessage.AuthMessage
	}{
		{
			name: "Test_New_Valid",
			args: struct {
				subID string
			}{
				subID: "test-subscription-id",
			},
			expected: &authmessage.AuthMessage{
				SubscriptionID: "test-subscription-id",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := authmessage.New(tt.args.subID)

			if authMessage.SubscriptionID != tt.expected.SubscriptionID {
				t.Errorf("authmessage.New() = %v, expected %v", authMessage.SubscriptionID, tt.expected.SubscriptionID)
			}
		})
	}
}

func TestAuthMessage_Marshal(t *testing.T) {
	tests := []struct {
		name     string
		message  *authmessage.AuthMessage
		expected []byte
	}{
		{
			name: "TestAuthMessage_Marshal_Valid",
			message: &authmessage.AuthMessage{
				SubscriptionID: "test-subscription-id",
			},
			expected: []byte(`["AUTH","test-subscription-id"]`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.message.Marshal()
			if err != nil {
				t.Errorf("authmessage.AuthMessage.Marshal() error = %v", err)
			}

			if !bytes.Equal(result, tt.expected) {
				t.Errorf("authmessage.AuthMessage.Marshal() = %v, expected %v", string(result), string(tt.expected))
			}
		})
	}
}
