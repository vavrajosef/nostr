package nostr

func NewRelay() *Relay {
	return &Relay{}
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
