package relay

type Limitations struct {
	MaxMessageLength uint `json:"max_message_length,omitempty"`
	MaxSubscriptions uint `json:"max_subscriptions,omitempty"`
	MaxFilters       uint `json:"max_filters,omitempty"`
	MaxLimit         uint `json:"max_limit,omitempty"`
	MaxSubIDLength   uint `json:"max_subid_length,omitempty"`
	MinPrefix        uint `json:"min_prefix,omitempty"`
	MaxEventTags     uint `json:"max_event_tags,omitempty"`
	MaxContentLength uint `json:"max_content_length,omitempty"`
	MinPowDifficulty uint `json:"min_pow_difficulty,omitempty"`
	AuthRequired     bool `json:"auth_required,omitempty"`
	PaymentRequired  bool `json:"payment_required,omitempty"`
}
