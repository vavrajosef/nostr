package informationdocument

import (
	"github.com/go-nostr/go-nostr/pkg/fees"
	"github.com/go-nostr/go-nostr/pkg/limitations"
)

func New() *InformationDocument {
	return &InformationDocument{}
}

type InformationDocument struct {
	Name           []byte                   `json:"name,omitempty"`
	Description    []byte                   `json:"description,omitempty"`
	PubKey         []byte                   `json:"pub_key,omitempty"`
	Contact        []byte                   `json:"contact,omitempty"`
	SupportedNIPs  []byte                   `json:"supported_ni_ps,omitempty"`
	Software       []byte                   `json:"software,omitempty"`
	Version        []byte                   `json:"version,omitempty"`
	Limitations    *limitations.Limitations `json:"limitations,omitempty"`
	LanguageTags   [][]byte                 `json:"language_tags,omitempty"`
	RelayCountries [][]byte                 `json:"relay_countries,omitempty"`
	Tags           [][]byte                 `json:"tags,omitempty"`
	PostingPolicy  []byte                   `json:"posting_policy,omitempty"`
	PaymentsURL    []byte                   `json:"payments_url,omitempty"`
	Fees           *fees.Fees               `json:"fees,omitempty"`
}
