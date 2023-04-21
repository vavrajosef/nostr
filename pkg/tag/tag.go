package tag

type Marker string

const (
	MarkerReply   Marker = "reply"
	MarkerRoot    Marker = "root"
	MarkerMention Marker = "mention"
)

// Type represents the type of a tag.
type Type string

const (
	TypeAmount      Type = "amount"          // TypeAmount represents millisats.
	TypeBadgeDesc   Type = "description"     // TypeBadgeDesc represents a badge description.
	TypeBadgeName   Type = "name"            // TypeBadgeName represents a badge name.
	TypeBolt11      Type = "bolt11"          // TypeBolt11 represents a Bolt11 invoice.
	TypeChallenge   Type = "challenge"       // TypeChallenge represents a challenge string.
	TypeContentWarn Type = "content-warning" // TypeContentWarn represents a content warning reason.
	TypeDelegation  Type = "delegation"      // TypeDelegation represents a delegation with pubkey, conditions, and delegation token.
	TypeEvent       Type = "a"               // TypeEvent represents coordinates to an event relay URL.
	TypeEventID     Type = "e"               // TypeEventID represents an event ID (hex) used in relay URLs and markers.
	TypeExpiration  Type = "expiration"      // TypeExpiration represents a Unix timestamp (string) for expiration.
	TypeGeohash     Type = "g"               // TypeGeohash represents a geohash.
	TypeHashtag     Type = "t"               // TypeHashtag represents a hashtag.
	TypeIdentifier  Type = "d"               // TypeIdentifier represents an identifier.
	TypeIdentity    Type = "i"               // TypeIdentity represents an identity used in proofs.
	TypeImage       Type = "image"           // TypeImage represents an image URL with dimensions in pixels.
	TypeInvDesc     Type = "description"     // TypeInvDesc represents an invoice description.
	TypeLNURL       Type = "lnurl"           // TypeLNURL represents a Bech32 encoded LNURL.
	TypeNonce       Type = "nonce"           // TypeNonce represents a random nonce.
	TypePreimage    Type = "preimage"        // TypePreimage represents a hash of a Bolt11 invoice.
	TypePubkey      Type = "p"               // TypePubkey represents a public key (hex) used in relay URLs.
	TypePublishedAt Type = "published_at"    // TypePublishedAt represents a Unix timestamp (string) for when an item was published.
	TypeReference   Type = "r"               // TypeReference represents a reference (URL, etc).
	TypeRelay       Type = "relay"           // TypeRelay represents a relay URL.
	TypeRelays      Type = "relays"          // TypeRelays represents a list of relays.
	TypeSubject     Type = "subject"         // TypeSubject represents a subject.
	TypeSummary     Type = "summary"         // TypeSummary represents an article summary.
	TypeThumb       Type = "thumb"           // TypeThumb represents a badge thumbnail with dimensions in pixels.
	TypeTitle       Type = "title"           // TypeTitle represents an article title.
	TypeZap         Type = "zap"             // TypeZap represents a profile name with a type of value.
)

// Tag is an interface for different types of tags.
type Tag interface {
	Marshal() []byte
}
