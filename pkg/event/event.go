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
	SetMetadata     Kind = 0 // Event for setting metadata
	TextNote        Kind = 1 // Event for storing a text note
	RecommendServer Kind = 2 // Event for recommending a server
)

// Options contains the necessary information for creating a new event.
type Options struct {
	Content    []byte                // Content to be included in the event
	Kind       Kind                  // The type of event
	PrivateKey *secp256k1.PrivateKey // The creator's private key
	Tags       []tag.Tag             // Tags associated with the event
}

// Event represents a single event object.
type Event struct {
	ID        []byte    `json:"id"`         // The unique identifier for the event
	PublicKey []byte    `json:"pk"`         // The creator's public key
	CreatedAt int64     `json:"created_at"` // The timestamp of event creation
	Kind      Kind      `json:"kind"`       // The type of event
	Tags      []tag.Tag `json:"tags"`       // Tags associated with the event
	Content   []byte    `json:"content"`    // The content of the event
	Signature []byte    `json:"sig"`        // The signature of the event
}

// New creates a new event with the provided options.
func New(opts *Options) (*Event, error) {
	// Encode the public key to a hex-encoded string
	pubKeyHex := make([]byte, hex.EncodedLen(len(opts.PrivateKey.PubKey().SerializeUncompressed())))
	hex.Encode(pubKeyHex, opts.PrivateKey.PubKey().SerializeCompressed())
	pubKey := bytes.ToLower(pubKeyHex)[:32]

	createdAt := time.Now().Unix()

	// Encode the tags
	encodedTags := make([][]byte, 0)
	for _, t := range opts.Tags {
		encodedTags = append(encodedTags, t.Encode())
	}
	tagsJSON, _ := json.Marshal(encodedTags)

	// Serialize the event data
	eventData := make([]byte, 0)
	eventData = append(eventData, 0)
	eventData = append(eventData, pubKey...)
	eventData = append(eventData, byte(opts.Kind))
	eventData = append(eventData, tagsJSON...)
	eventData = append(eventData, opts.Content...)

	// Encode the event data to a hex-encoded string
	eventDataJSON, err := json.Marshal(eventData)
	if err != nil {
		return nil, err
	}
	eventDataHex := make([]byte, hex.EncodedLen(len(eventDataJSON)))
	hex.Encode(eventDataHex, eventDataJSON)

	// Get the unique identifier for the event
	id := bytes.ToLower(eventDataHex)[:32]

	// Encrypt the content
	encryptedContent, err := secp256k1.Encrypt(opts.PrivateKey.PubKey(), opts.Content)
	if err != nil {
		return nil, err
	}

	// Create and return the new event
	return &Event{
		Content:   encryptedContent,
		CreatedAt: createdAt,
		ID:        id,
		Kind:      opts.Kind,
		PublicKey: pubKey,
		Tags:      opts.Tags,
	}, nil
}
