package nostr_test

import (
	"testing"

	nostr "github.com/go-nostr/go-nostr/pkg"
)

func Test_NewRelay(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			opt *nostr.RelayOptions
		}
		expected *nostr.Relay
	}{
		{
			name: "New Relay",
			args: struct {
				opt *nostr.RelayOptions
			}{
				opt: &nostr.RelayOptions{},
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			relay, err := nostr.NewRelay(tt.args.opt)
			if err != nil {
				t.Errorf("event.New() error = %v", err)
			}

			t.Logf("%v", relay)
		})
	}
}

func TestRelay_Dial(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			opt *nostr.RelayOptions
		}
		expected *nostr.Relay
	}{
		// TODO: temporarily disabled
		// {
		// 	name: "Relay dial",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			relay, err := nostr.NewRelay(tt.args.opt)
			if err != nil {
				t.Errorf("event.New() error = %v", err)
			}

			relay.Dial()

			t.Logf("%v", relay)
		})
	}
}
