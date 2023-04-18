package relay

type Admission struct {
	Kinds  []uint `json:"kinds,omitempty"`
	Amount uint   `json:"amount,omitempty"`
	Unit   []byte `json:"unit,omitempty"`
}
