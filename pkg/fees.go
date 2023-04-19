package nostr

func NewFees() *Fees {
	return &Fees{}
}

type Fees struct {
	Admission *Admission `json:"admission,omitempty"`
}
