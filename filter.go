package nostr

// Filter is a struct that defines a set of criteria for filtering events.
type Filter struct {
	IDs        []string    `json:"ids,omitempty"`     // The IDs of the events to filter
	Authors    []string    `json:"authors,omitempty"` // The authors of the events to filter
	Kinds      []EventKind `json:"kinds,omitempty"`   // The kinds of the events to filter
	EventIDs   []string    `json:"#e,omitempty"`      // The event IDs to filter by
	PublicKeys []string    `json:"#p,omitempty"`      // The public keys to filter by
	Since      int64       `json:"since,omitempty"`   // The starting timestamp for filtering events
	Until      int64       `json:"until,omitempty"`   // The ending timestamp for filtering events
	Limit      uint        `json:"limit,omitempty"`   // The maximum number of events to return
	Search     string      `json:"search,omitempty"`  // TBD
}
