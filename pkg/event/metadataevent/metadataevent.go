package metadataevent

import "github.com/go-nostr/go-nostr/pkg/event"

type MetadataEventOptions struct {
	*event.Options
}

type MetadataEvent struct {
	*event.Event
}

func New(opt *MetadataEventOptions) (*MetadataEvent, error) {
	evt, err := event.New(&event.Options{
		Kind: event.KindMetadata,
	})
	if err != nil {
		return nil, err
	}

	return &MetadataEvent{evt}, nil
}
