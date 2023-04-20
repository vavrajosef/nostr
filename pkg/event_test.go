package nostr_test

import (
	"reflect"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1"
	nostr "github.com/go-nostr/go-nostr/pkg"
)

func Test_NewEvent(t *testing.T) {
	privateKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	tests := []struct {
		name string
		args struct {
			opt *nostr.EventOptions
		}
		expected *nostr.Event
	}{
		{
			name: "New Event",
			args: struct {
				opt *nostr.EventOptions
			}{
				opt: &nostr.EventOptions{
					Content:    []byte("This is a test message."),
					Kind:       nostr.EventKindShortTextNote,
					PrivateKey: privateKey,

					Tags: []nostr.Tag{},
				},
			},
			// TODO: replace
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evt, err := nostr.NewEvent(tt.args.opt)
			if err != nil {
				t.Errorf("event.New() error = %v", err)
			}

			// TODO: replace
			// if !reflect.DeepEqual(evt.Content, tt.expected) {
			// 	t.Errorf("event.New() = %v, expected = %v", evt.Content, tt.expected)
			// }

			if evt.Kind != tt.args.opt.Kind {
				t.Errorf("event.NewEvent() Kind = %v, expected = %v", evt.Kind, tt.args.opt.Kind)
			}

			if !reflect.DeepEqual(evt.Tags, tt.args.opt.Tags) {
				t.Errorf("event.NewEvent() Tags = %v, expected = %v", evt.Tags, tt.args.opt.Tags)
			}

			t.Logf("%v", evt)
		})
	}
}
