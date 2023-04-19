package nostr

func NewCommandResult() *CommandResult {
	return &CommandResult{}
}

type CommandResult struct {
	Status  []byte `json:"status,omitempty"`
	EventID []byte `json:"event_id,omitempty"`
	IsSaved bool   `json:"is_saved,omitempty"`
	Message []byte `json:"message,omitempty"`
}
