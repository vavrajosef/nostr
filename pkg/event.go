package nostr

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/decred/dcrd/dcrec/secp256k1"
)

// Kind represents the different types of events available.
type EventKind uint32

const (
	EventKindMetadata                EventKind = 0     // Event for setting metadata
	EventKindShortTextNote           EventKind = 1     // Event for storing a text note
	EventKindRecommendRelay          EventKind = 2     // Event for recommending a server
	EventKindContacts                EventKind = 3     // Event for contact List and petnames
	EventKindEncryptedDirectMessages EventKind = 4     // Event for encrypted direct messages
	EventKindEventDeletion           EventKind = 5     // Event for deleting events
	EventKindReposts                 EventKind = 6     // Event for signaling to followers that another event is worth reading
	EventKindReaction                EventKind = 7     // Event for reacting to other notes
	EventKindBadgeAward              EventKind = 8     // Event for awarding badges to users
	EventKindChannelCreation         EventKind = 40    // Event for creating new channels
	EventKindChannelMetadata         EventKind = 41    // Event for setting channel metadata
	EventKindChannelMessage          EventKind = 42    // Event for posting messages in a channel
	EventKindChannelHideMessage      EventKind = 43    // Event for hiding messages in a channel
	EventKindChannelMuteUser         EventKind = 44    // Event for muting a user in a channel
	EventKindReporting               EventKind = 1984  // Event for reporting content or users
	EventKindZapRequest              EventKind = 9734  // Event for requesting a Zap action
	EventKindZap                     EventKind = 9735  // Event for performing a Zap action
	EventKindMuteList                EventKind = 10000 // Event for managing a mute list
	EventKindPinList                 EventKind = 10001 // Event for managing a pin list
	EventKindRelayListMetadata       EventKind = 10002 // Event for managing relay list metadata
	EventKindClientAuthentication    EventKind = 22242 // Event for client authentication process
	EventKindNostrConnect            EventKind = 24133 // Event for Nostr client connection
	EventKindCategorizedPeopleList   EventKind = 30000 // Event for managing categorized people list
	EventKindCategorizedBookmarkList EventKind = 30001 // Event for managing categorized bookmark list
	EventKindProfileBadges           EventKind = 30008 // Event for profile badges management
	EventKindBadgeDefinition         EventKind = 30009 // Event for defining badges
	EventKindCreateOrUpdateStall     EventKind = 30017 // Event for creating or updating a stall
	EventKindCreateOrUpdateProduct   EventKind = 30018 // Event for creating or updating a product
	EventKindLongFormContent         EventKind = 30023 // Event for posting long-form content
	EventKindApplicationSpecificData EventKind = 30078 // Event for managing application-specific data
)

// EventOptions contains the necessary information for creating a new event.
type EventOptions struct {
	Content    []byte                `json:"content,omitempty"`     // Content to be included in the event
	Kind       EventKind             `json:"kind,omitempty"`        // The type of event
	OTS        []byte                `json:"ots,omitempty"`         // TBD
	PrivateKey *secp256k1.PrivateKey `json:"private_key,omitempty"` // The creator's private key
	Tags       []Tag                 `json:"tags,omitempty"`        // Tags associated with the event
}

// Event represents a single event object.
type Event struct {
	Content   []byte    `json:"content,omitempty"`    // The content of the event
	CreatedAt int64     `json:"created_at,omitempty"` // The timestamp of event creation
	ID        []byte    `json:"id,omitempty"`         // The unique identifier for the event
	Kind      EventKind `json:"kind,omitempty"`       // The type of event
	OTS       []byte    `json:"ots,omitempty"`        // TBD
	PublicKey []byte    `json:"pk,omitempty"`         // The creator's public key
	Signature []byte    `json:"sig,omitempty"`        // The signature of the event
	Tags      []Tag     `json:"tags,omitempty"`       // Tags associated with the event
}

func (e *Event) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// NewEvent creates a new event with the provided options.
func NewEvent(opts *EventOptions) (*Event, error) {
	// NOTE: TBD
	pubKeyHex := make([]byte, hex.EncodedLen(len(opts.PrivateKey.PubKey().SerializeUncompressed())))

	// NOTE: TBD
	hex.Encode(pubKeyHex, opts.PrivateKey.PubKey().SerializeCompressed())

	// NOTE: TBD
	pubKey := bytes.ToLower(pubKeyHex)[:32]

	createdAt := time.Now().Unix()

	// NOTE: TBD
	encodedTags := make([][]byte, 0)

	// NOTE: TBD
	for _, t := range opts.Tags {
		encodedTags = append(encodedTags, t.Marshal())
	}

	// NOTE: TBD
	tagsJSON, _ := json.Marshal(encodedTags)

	// NOTE: TBD
	eventData := make([]byte, 0)
	eventData = append(eventData, 0)
	eventData = append(eventData, pubKey...)
	eventData = append(eventData, byte(opts.Kind))
	eventData = append(eventData, tagsJSON...)
	eventData = append(eventData, opts.Content...)

	// NOTE: TBD
	eventDataJSON, err := json.Marshal(eventData)
	if err != nil {
		return nil, err
	}

	// NOTE: TBD
	eventDataHex := make([]byte, hex.EncodedLen(len(eventDataJSON)))

	// NOTE: TBD
	hex.Encode(eventDataHex, eventDataJSON)

	// NOTE: TBD
	// NOTE: TBD
	id := bytes.ToLower(eventDataHex)[:32]

	// NOTE: TBD
	encryptedContent, err := secp256k1.Encrypt(opts.PrivateKey.PubKey(), opts.Content)
	if err != nil {
		return nil, err
	}

	// NOTE: TBD
	evt := &Event{
		Content:   encryptedContent,
		CreatedAt: createdAt,
		ID:        id,
		Kind:      opts.Kind,
		PublicKey: pubKey,
		Tags:      opts.Tags,
	}

	// NOTE: TBD
	return evt, nil
}

type MetadataEventOptions struct {
	*EventOptions
}

type MetadataEvent struct {
	*Event
}

func NewMetadataEvent(opt *MetadataEventOptions) (*MetadataEvent, error) {
	evt, err := NewEvent(&EventOptions{
		Kind: EventKindMetadata,
	})
	if err != nil {
		return nil, err
	}

	return &MetadataEvent{evt}, nil
}
