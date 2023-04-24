package nostr

import "encoding/json"

type TagType string

const (
	TagTypeAmount      TagType = "amount"          // TagTypeAmount represents millisats.
	TagTypeBadgeDesc   TagType = "description"     // TagTypeBadgeDesc represents a badge description.
	TagTypeBadgeName   TagType = "name"            // TagTypeBadgeName represents a badge name.
	TagTypeBolt11      TagType = "bolt11"          // TagTypeBolt11 represents a Bolt11 invoice.
	TagTypeChallenge   TagType = "challenge"       // TagTypeChallenge represents a challenge string.
	TagTypeContentWarn TagType = "content-warning" // TagTypeContentWarn represents a content warning reason.
	TagTypeDelegation  TagType = "delegation"      // TagTypeDelegation represents a delegation with pubkey, conditions, and delegation token.
	TagTypeEvent       TagType = "a"               // TagTypeEvent represents coordinates to an event relay URL.
	TagTypeEventID     TagType = "e"               // TagTypeEventID represents an event ID (hex) used in relay URLs and markers.
	TagTypeExpiration  TagType = "expiration"      // TagTypeExpiration represents a Unix timestamp (string) for expiration.
	TagTypeGeohash     TagType = "g"               // TagTypeGeohash represents a geohash.
	TagTypeHashtag     TagType = "t"               // TagTypeHashtag represents a hashtag.
	TagTypeIdentifier  TagType = "d"               // TagTypeIdentifier represents an identifier.
	TagTypeIdentity    TagType = "i"               // TagTypeIdentity represents an identity used in proofs.
	TagTypeImage       TagType = "image"           // TagTypeImage represents an image URL with dimensions in pixels.
	TagTypeInvDesc     TagType = "description"     // TagTypeInvDesc represents an invoice description.
	TagTypeLNURL       TagType = "lnurl"           // TagTypeLNURL represents a Bech32 encoded LNURL.
	TagTypeNonce       TagType = "nonce"           // TagTypeNonce represents a random nonce.
	TagTypePreimage    TagType = "preimage"        // TagTypePreimage represents a hash of a Bolt11 invoice.
	TagTypePubkey      TagType = "p"               // TagTypePubkey represents a public key (hex) used in relay URLs.
	TagTypePublishedAt TagType = "published_at"    // TagTypePublishedAt represents a Unix timestamp (string) for when an item was published.
	TagTypeReference   TagType = "r"               // TagTypeReference represents a reference (URL, etc).
	TagTypeRelay       TagType = "relay"           // TagTypeRelay represents a relay URL.
	TagTypeRelays      TagType = "relays"          // TagTypeRelays represents a list of relays.
	TagTypeSubject     TagType = "subject"         // TagTypeSubject represents a subject.
	TagTypeSummary     TagType = "summary"         // TagTypeSummary represents an article summary.
	TagTypeThumb       TagType = "thumb"           // TagTypeThumb represents a badge thumbnail with dimensions in pixels.
	TagTypeTitle       TagType = "title"           // TagTypeTitle represents an article title.
	TagTypeZap         TagType = "zap"             // TagTypeZap represents a profile name with a type of value.
)

// Tag is an interface for different types of tags.
type Tag interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

// NewAmountTag creates a new AmountTag with the provided amount.
func NewAmountTag(amount uint32) Tag {
	return &AmountTag{
		TagTypeAmount,
		amount,
	}
}

// AmountTag TBD
type AmountTag struct {
	Type   TagType `json:"type,omitempty"`
	Amount uint32  `json:"amount,omitempty"`
}

// Marshal TBD
func (t *AmountTag) Marshal() ([]byte, error) {
	return json.Marshal([]interface{}{
		t.Type,
		t.Amount,
	})
}

// Unmarshal TBD
func (t *AmountTag) Unmarshal(data []byte) error {
	args := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &t.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(args[1], &t.Amount); err != nil {
		return err
	}
	return nil
}

// TODO: Add other tags or take different approach
