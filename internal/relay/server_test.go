package relay_test

import (
	"testing"

	"github.com/go-nostr/nostr/internal/relay"
)

func TestRelay_Dial(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			relayServer := relay.NewServer()

			relayServer.Dial()
		})
	}
}
