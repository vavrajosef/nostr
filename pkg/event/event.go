package event

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/go-nostr/go-nostr/pkg/tag"
)

// Kind represents the different types of events available.
type Kind uint32

const (
	KindMetadata                Kind = 0     // Event for setting metadata
	KindShortTextNote           Kind = 1     // Event for storing a text note
	KindRecommendRelay          Kind = 2     // Event for recommending a server
	KindContacts                Kind = 3     // Event for contact List and petnames
	KindEncryptedDirectMessages Kind = 4     // Event for encrypted direct messages
	KindEventDeletion           Kind = 5     // Event for deleting events
	KindReposts                 Kind = 6     // Event for signaling to followers that another event is worth reading
	KindReaction                Kind = 7     // Event for reacting to other notes
	KindBadgeAward              Kind = 8     // Event for awarding badges to users
	KindChannelCreation         Kind = 40    // Event for creating new channels
	KindChannelMetadata         Kind = 41    // Event for setting channel metadata
	KindChannelMessage          Kind = 42    // Event for posting messages in a channel
	KindChannelHideMessage      Kind = 43    // Event for hiding messages in a channel
	KindChannelMuteUser         Kind = 44    // Event for muting a user in a channel
	KindReporting               Kind = 1984  // Event for reporting content or users
	KindZapRequest              Kind = 9734  // Event for requesting a Zap action
	KindZap                     Kind = 9735  // Event for performing a Zap action
	KindMuteList                Kind = 10000 // Event for managing a mute list
	KindPinList                 Kind = 10001 // Event for managing a pin list
	KindRelayListMetadata       Kind = 10002 // Event for managing relay list metadata
	KindClientAuthentication    Kind = 22242 // Event for client authentication process
	KindNostrConnect            Kind = 24133 // Event for Nostr client connection
	KindCategorizedPeopleList   Kind = 30000 // Event for managing categorized people list
	KindCategorizedBookmarkList Kind = 30001 // Event for managing categorized bookmark list
	KindProfileBadges           Kind = 30008 // Event for profile badges management
	KindBadgeDefinition         Kind = 30009 // Event for defining badges
	KindCreateOrUpdateStall     Kind = 30017 // Event for creating or updating a stall
	KindCreateOrUpdateProduct   Kind = 30018 // Event for creating or updating a product
	KindLongFormContent         Kind = 30023 // Event for posting long-form content
	KindApplicationSpecificData Kind = 30078 // Event for managing application-specific data
)

// Options contains the necessary information for creating a new event.
type Options struct {
	Content    []byte                `json:"content,omitempty"`     // Content to be included in the event
	Kind       Kind                  `json:"kind,omitempty"`        // The type of event
	OTS        []byte                `json:"ots,omitempty"`         // TBD
	PrivateKey *secp256k1.PrivateKey `json:"private_key,omitempty"` // The creator's private key
	Tags       []tag.Tag             `json:"tags,omitempty"`        // Tags associated with the event
}

// Event represents a single event object.
type Event struct {
	Content   []byte    `json:"content,omitempty"`    // The content of the event
	CreatedAt int64     `json:"created_at,omitempty"` // The timestamp of event creation
	ID        []byte    `json:"id,omitempty"`         // The unique identifier for the event
	Kind      Kind      `json:"kind,omitempty"`       // The type of event
	OTS       []byte    `json:"ots,omitempty"`        // TBD
	PublicKey []byte    `json:"pk,omitempty"`         // The creator's public key
	Signature []byte    `json:"sig,omitempty"`        // The signature of the event
	Tags      []tag.Tag `json:"tags,omitempty"`       // Tags associated with the event
}

func (e *Event) Marshal() []byte {
	b, _ := json.Marshal(e)

	return b
}

// New creates a new event with the provided options.
func New(opts *Options) (*Event, error) {
	// NOTE: Encode the public key to a hex-encoded string
	pubKeyHex := make([]byte, hex.EncodedLen(len(opts.PrivateKey.PubKey().SerializeUncompressed())))
	hex.Encode(pubKeyHex, opts.PrivateKey.PubKey().SerializeCompressed())
	pubKey := bytes.ToLower(pubKeyHex)[:32]

	createdAt := time.Now().Unix()

	// NOTE: Encode the tags
	encodedTags := make([][]byte, 0)
	for _, t := range opts.Tags {
		encodedTags = append(encodedTags, t.Marshal())
	}
	tagsJSON, _ := json.Marshal(encodedTags)

	// NOTE: Serialize the event data
	eventData := make([]byte, 0)
	eventData = append(eventData, 0)
	eventData = append(eventData, pubKey...)
	eventData = append(eventData, byte(opts.Kind))
	eventData = append(eventData, tagsJSON...)
	eventData = append(eventData, opts.Content...)

	// NOTE: Encode the event data to a hex-encoded string
	eventDataJSON, err := json.Marshal(eventData)
	if err != nil {
		return nil, err
	}
	eventDataHex := make([]byte, hex.EncodedLen(len(eventDataJSON)))
	hex.Encode(eventDataHex, eventDataJSON)

	// NOTE: Get the unique identifier for the event
	id := bytes.ToLower(eventDataHex)[:32]

	// NOTE: Encrypt the content
	encryptedContent, err := secp256k1.Encrypt(opts.PrivateKey.PubKey(), opts.Content)
	if err != nil {
		return nil, err
	}

	// NOTE: Create and return the new event
	return &Event{
		Content:   encryptedContent,
		CreatedAt: createdAt,
		ID:        id,
		Kind:      opts.Kind,
		PublicKey: pubKey,
		Tags:      opts.Tags,
	}, nil
}
