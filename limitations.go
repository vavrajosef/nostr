package nostr

// Limitations TBD
type Limitations struct {
	MaxMessageLength int64 `json:"max_message_length,omitempty"`
	MaxSubscriptions int64 `json:"max_subscriptions,omitempty"`
	MaxFilters       int64 `json:"max_filters,omitempty"`
	MaxLimit         int64 `json:"max_limit,omitempty"`
	MaxSubIDLength   int64 `json:"max_subid_length,omitempty"`
	MinPrefix        int64 `json:"min_prefix,omitempty"`
	MaxEventTags     int64 `json:"max_event_tags,omitempty"`
	MaxContentLength int64 `json:"max_content_length,omitempty"`
	MinPowDifficulty int64 `json:"min_pow_difficulty,omitempty"`
	AuthRequired     bool  `json:"auth_required,omitempty"`
	PaymentRequired  bool  `json:"payment_required,omitempty"`
}
