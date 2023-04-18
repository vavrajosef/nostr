package event_test

import (
	"reflect"
	"testing"

	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/go-nostr/go-nostr/pkg/event"
	"github.com/go-nostr/go-nostr/pkg/tag"
)

func Test_New(t *testing.T) {
	privateKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	tests := []struct {
		name string
		args struct {
			opt *event.Options
		}
		expected *event.Event
	}{
		{
			name: "New Event",
			args: struct {
				opt *event.Options
			}{
				opt: &event.Options{
					Content:    []byte("This is a test message."),
					Kind:       event.TextNote,
					PrivateKey: privateKey,

					Tags: []tag.Tag{},
				},
			},
			expected: nil, // We will compare the values inside the test, so we keep this as nil.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evt, err := event.New(tt.args.opt)
			if err != nil {
				t.Errorf("event.New() error = %v", err)
			}

			// if !reflect.DeepEqual(evt.Content, tt.expected) {
			// 	t.Errorf("event.New() = %v, expected = %v", evt.Content, tt.expected)
			// }

			if evt.Kind != tt.args.opt.Kind {
				t.Errorf("event.New() Kind = %v, expected = %v", evt.Kind, tt.args.opt.Kind)
			}

			if !reflect.DeepEqual(evt.Tags, tt.args.opt.Tags) {
				t.Errorf("event.New() Tags = %v, expected = %v", evt.Tags, tt.args.opt.Tags)
			}

			t.Logf("%v", evt)
		})
	}
}
