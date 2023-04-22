package nostr

// Admission ...
type Admission struct {
	Kinds  EventKind `json:"kinds,omitempty"`
	Amount uint      `json:"amount,omitempty"`
	Unit   string    `json:"unit,omitempty"`
}
