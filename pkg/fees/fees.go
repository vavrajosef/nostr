package fees

import "github.com/go-nostr/go-nostr/pkg/admission"

func New(opt *Options) *Fees {
	return &Fees{opt}
}

type Options struct {
	Admission *admission.Admission `json:"admission,omitempty"`
}

type Fees struct {
	*Options
}
