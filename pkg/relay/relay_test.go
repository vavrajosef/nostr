package relay_test

import (
	"testing"

	"github.com/go-nostr/go-nostr/pkg/relay"
)

func Test_NewRelay(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			opt *relay.Options
		}
		expected *relay.Relay
	}{
		{
			name: "New Relay",
			args: struct {
				opt *relay.Options
			}{
				opt: &relay.Options{},
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			relay, err := relay.NewRelay(tt.args.opt)
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
			opt *relay.Options
		}
		expected *relay.Relay
	}{
		// TODO: temporarily disabled
		// {
		// 	name: "Relay dial",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			relay, err := relay.NewRelay(tt.args.opt)
			if err != nil {
				t.Errorf("event.New() error = %v", err)
			}

			relay.Dial()

			t.Logf("%v", relay)
		})
	}
}
