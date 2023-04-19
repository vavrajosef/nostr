package nostr

func NewAdmission() *Admission {
	return &Admission{}
}

type Admission struct {
	Kinds  []uint `json:"kinds,omitempty"`
	Amount uint   `json:"amount,omitempty"`
	Unit   []byte `json:"unit,omitempty"`
}
