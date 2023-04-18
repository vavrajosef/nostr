package relay

func New(opt *Options) *Relay {
	return &Relay{}
}

type Admission struct {
	Kinds  []uint `json:"kinds,omitempty"`
	Amount uint   `json:"amount,omitempty"`
	Unit   []byte `json:"unit,omitempty"`
}

type Fees struct {
	Admission *Admission `json:"admission,omitempty"`
}

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

type Options struct {
	Name          []byte       `json:"name,omitempty"`
	Description   []byte       `json:"description,omitempty"`
	PubKey        []byte       `json:"pub_key,omitempty"`
	Contact       []byte       `json:"contact,omitempty"`
	SupportedNIPs []byte       `json:"supported_ni_ps,omitempty"`
	Software      []byte       `json:"software,omitempty"`
	Version       []byte       `json:"version,omitempty"`
	Limitations   *Limitations `json:"limitations,omitempty"`
}

type Relay struct {
	Name          []byte       `json:"name,omitempty"`
	Description   []byte       `json:"description,omitempty"`
	PubKey        []byte       `json:"pub_key,omitempty"`
	Contact       []byte       `json:"contact,omitempty"`
	SupportedNIPs []byte       `json:"supported_ni_ps,omitempty"`
	Software      []byte       `json:"software,omitempty"`
	Version       []byte       `json:"version,omitempty"`
	Limitations   *Limitations `json:"limitations,omitempty"`
}

type RelayInformationDocument struct {
	Name           []byte       `json:"name,omitempty"`
	Description    []byte       `json:"description,omitempty"`
	PubKey         []byte       `json:"pub_key,omitempty"`
	Contact        []byte       `json:"contact,omitempty"`
	SupportedNIPs  []byte       `json:"supported_ni_ps,omitempty"`
	Software       []byte       `json:"software,omitempty"`
	Version        []byte       `json:"version,omitempty"`
	Limitations    *Limitations `json:"limitations,omitempty"`
	LanguageTags   [][]byte     `json:"language_tags,omitempty"`
	RelayCountries [][]byte     `json:"relay_countries,omitempty"`
	Tags           [][]byte     `json:"tags,omitempty"`
	PostingPolicy  []byte       `json:"posting_policy,omitempty"`
	PaymentsURL    []byte       `json:"payments_url,omitempty"`
	Fees           *Fees        `json:"fees,omitempty"`
}

type CommandResult struct {
	Status  []byte `json:"status,omitempty"`
	EventID []byte `json:"event_id,omitempty"`
	IsSaved bool   `json:"is_saved,omitempty"`
	Message []byte `json:"message,omitempty"`
}
