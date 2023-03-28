package event

import (
	"encoding/base64"
	"time"
)

// https://github.com/nostr-protocol/nips/blob/master/01.md#basic-event-kinds
const (
	EventKindSetMetadata       = 0
	EventKindTextNote          = 1
	EventKindRecommendedServer = 2
)

type EventOptions struct {
	Content   string
	EventData []byte
	Kind      uint32
	Tags      [][]byte
}

func New(opt *EventOptions) (*Event, error) {
	// TODO: replace with secp256k1
	id := make([]byte, 32)
	publicKey := make([]byte, 32)
	createdAt := time.Now()

	base64.StdEncoding.Encode(id, opt.EventData)

	return &Event{
		Content:   opt.Content,
		CreatedAt: createdAt,
		ID:        id,
		Kind:      opt.Kind,
		PublicKey: publicKey,
		Tags:      opt.Tags,
	}, nil
}

func Validate(evt *Event) error {
	return nil
}

type Event struct {
	ID        []byte
	PublicKey []byte
	CreatedAt time.Time
	Kind      uint32
	Tags      [][]byte
	Content   string
	Sig       string
}
