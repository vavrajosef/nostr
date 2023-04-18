package filter

// New creates a new Filter instance using the provided Options.
func New(opt *Options) *Filter {
	return &Filter{
		IDs:        opt.IDs,
		Authors:    opt.Authors,
		Kinds:      opt.Kinds,
		EventIDs:   opt.EventIDs,
		PublicKeys: opt.PublicKeys,
		Since:      opt.Since,
		Until:      opt.Until,
		Limit:      opt.Limit,
		Search:     opt.Search,
	}
}

// Options is a struct containing options for creating a new Filter.
type Options struct {
	IDs        [][]byte `json:"ids,omitempty"`     // The IDs of the events to filter
	Authors    [][]byte `json:"authors,omitempty"` // The authors of the events to filter
	Kinds      [][]byte `json:"kinds,omitempty"`   // The kinds of the events to filter
	EventIDs   [][]byte `json:"#e,omitempty"`      // The event IDs to filter by
	PublicKeys [][]byte `json:"#p,omitempty"`      // The public keys to filter by
	Since      string   `json:"since,omitempty"`   // The starting timestamp for filtering events
	Until      string   `json:"until,omitempty"`   // The ending timestamp for filtering events
	Limit      uint     `json:"limit,omitempty"`   // The maximum number of events to return
	Search     []byte   `json:"search,omitempty"`  // TBD
}

// Filter is a struct that defines a set of criteria for filtering events.
type Filter struct {
	IDs        [][]byte `json:"ids,omitempty"`     // The IDs of the events to filter
	Authors    [][]byte `json:"authors,omitempty"` // The authors of the events to filter
	Kinds      [][]byte `json:"kinds,omitempty"`   // The kinds of the events to filter
	EventIDs   [][]byte `json:"#e,omitempty"`      // The event IDs to filter by
	PublicKeys [][]byte `json:"#p,omitempty"`      // The public keys to filter by
	Since      string   `json:"since,omitempty"`   // The starting timestamp for filtering events
	Until      string   `json:"until,omitempty"`   // The ending timestamp for filtering events
	Limit      uint     `json:"limit,omitempty"`   // The maximum number of events to return
	Search     []byte   `json:"search,omitempty"`  // TBD
}
