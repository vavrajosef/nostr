package nostr

// CommandResult TBD
type CommandResult struct {
	Status  string `json:"status,omitempty"`
	EventID string `json:"event_id,omitempty"`
	IsSaved bool   `json:"is_saved,omitempty"`
	Message string `json:"message,omitempty"`
}
