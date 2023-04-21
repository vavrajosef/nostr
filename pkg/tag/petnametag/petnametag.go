package identifiertag

import (
	"encoding/json"

	"github.com/go-nostr/go-nostr/pkg/tag"
)

// New TBD
func New(pubKey []byte, relayURL []byte) *PetnameTag {
	return &PetnameTag{
		Type:     tag.TypePetName,
		Petname:  pubKey,
		RelayURL: relayURL,
	}
}

// PetnameTag is a tag for a public key, including the public key and relay URL.
type PetnameTag struct {
	Type     tag.Type `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte   `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte   `json:"relay_url,omitempty"` // The URL of the relay server for the public key
	Petname  []byte   `json:"petname,omitempty"`   // TBD
}

// Encode encodes the PetnameTag into a byte slice.
func (t *PetnameTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(tag.TypePetName))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)
	b = append(b, t.Petname)

	return b
}

func (t *PetnameTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}
