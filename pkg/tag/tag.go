package tag

import "encoding/json"

type Marker string

const (
	MarkerReply   Marker = "reply"
	MarkerRoot    Marker = "root"
	MarkerMention Marker = "mention"
)

// Type represents the type of a tag.
type Type string

// Tag is an interface for different types of tags.
type Tag interface {
	Encode() [][]byte
	Marshal() []byte
}

// TODO: add/change types
const (
	TypeA              Type = "a"
	TypeD              Type = "d"
	TypeEvent          Type = "e"
	TypePubKey         Type = "p"
	TypeNonce          Type = "nonce"
	TypeSubject        Type = "subject"
	TypeContentWarning Type = "content-warning"
	TypeIdentity       Type = "i"
	TypeExpiration     Type = "expiration"
)

// NewEventTag TBD
func NewEventTag(eventID []byte, relayURL []byte) *EventTag {
	return &EventTag{
		Type:     TypeEvent,
		EventID:  eventID,
		RelayURL: relayURL,
	}
}

// EventTag is a tag for an event, including the event ID and relay URL.
type EventTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "e" for event tags
	EventID  []byte `json:"event_id,omitempty"`  // The unique identifier for the associated event
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the event
	Marker   []byte `json:"marker,omitempty"`    // The URL of the relay server for the event
}

// Encode encodes the EventTag into a byte slice.
func (t *EventTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypeEvent))
	b = append(b, t.EventID)
	b = append(b, t.RelayURL)

	return b
}

func (t *EventTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewPetnameTag TBD
func NewPetnameTag(pubKey []byte, relayURL []byte) *PetnameTag {
	return &PetnameTag{
		Type:     TypeEvent,
		Petname:  pubKey,
		RelayURL: relayURL,
	}
}

// PetnameTag is a tag for a public key, including the public key and relay URL.
type PetnameTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
	Petname  []byte `json:"petname,omitempty"`   // TBD
}

// Encode encodes the PetnameTag into a byte slice.
func (t *PetnameTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypePubKey))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)
	b = append(b, t.Petname)

	return b
}

func (t *PetnameTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewPubKeyTag TBD
func NewPubKeyTag(pubKey []byte, relayURL []byte) *PubKeyTag {
	return &PubKeyTag{
		Type:     TypeEvent,
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// PubKeyTag is a tag for a public key, including the public key and relay URL.
type PubKeyTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the PubKeyTag into a byte slice.
func (t *PubKeyTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypePubKey))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *PubKeyTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewQueryTag TBD
func NewQueryTag(queryType []byte, pubKey []byte, relayURL []byte) *QueryTag {
	return &QueryTag{
		Type:     Type(queryType),
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// QueryTag is a tag for a public key, including the public key and relay URL.
type QueryTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the QueryTag into a byte slice.
func (t *QueryTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypePubKey))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *QueryTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewNonceTag TBD
func NewNonceTag(pubKey []byte, relayURL []byte) *NonceTag {
	return &NonceTag{
		Type:     TypeEvent,
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// NonceTag is a tag for a public key, including the public key and relay URL.
type NonceTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the NonceTag into a byte slice.
func (t *NonceTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypePubKey))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *NonceTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewSubjectTag TBD
func NewSubjectTag(pubKey []byte, relayURL []byte) *SubjectTag {
	return &SubjectTag{
		Type:     TypeNonce,
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// SubjectTag is a tag for a public key, including the public key and relay URL.
type SubjectTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the SubjectTag into a byte slice.
func (t *SubjectTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypeSubject))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *SubjectTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewDelegationTag TBD
func NewDelegationTag(pubKey []byte, relayURL []byte) *DelegationTag {
	return &DelegationTag{
		Type:     TypeNonce,
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// DelegationTag is a tag for a public key, including the public key and relay URL.
type DelegationTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the DelegationTag into a byte slice.
func (t *DelegationTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypeSubject))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *DelegationTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewContentWarningTag TBD
func NewContentWarningTag(pubKey []byte, relayURL []byte) *ContentWarningTag {
	return &ContentWarningTag{
		Type:     TypeNonce,
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// ContentWarningTag is a tag for a public key, including the public key and relay URL.
type ContentWarningTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the ContentWarningTag into a byte slice.
func (t *ContentWarningTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypeSubject))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *ContentWarningTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewIdentityTag TBD
func NewIdentityTag(pubKey []byte, relayURL []byte) *IdentityTag {
	return &IdentityTag{
		Type:     TypeNonce,
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// IdentityTag is a tag for a public key, including the public key and relay URL.
type IdentityTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the IdentityTag into a byte slice.
func (t *IdentityTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypeSubject))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *IdentityTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}

// NewExpirationTag TBD
func NewExpirationTag(pubKey []byte, relayURL []byte) *ExpirationTag {
	return &ExpirationTag{
		Type:     TypeNonce,
		PubKey:   pubKey,
		RelayURL: relayURL,
	}
}

// ExpirationTag is a tag for a public key, including the public key and relay URL.
type ExpirationTag struct {
	Type     Type   `json:"type,omitempty"`      // The type of the tag, which is "p" for public key tags
	PubKey   []byte `json:"pub_key,omitempty"`   // The public key associated with the tag
	RelayURL []byte `json:"relay_url,omitempty"` // The URL of the relay server for the public key
}

// Encode encodes the ExpirationTag into a byte slice.
func (t *ExpirationTag) Encode() [][]byte {
	b := make([][]byte, 0)
	b = append(b, []byte(TypeSubject))
	b = append(b, t.PubKey)
	b = append(b, t.RelayURL)

	return b
}

func (t *ExpirationTag) Marshal() []byte {
	b, _ := json.Marshal(t.Encode())

	return b
}
