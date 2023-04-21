package nostr

type Marker string

const (
	MarkerReply   Marker = "reply"
	MarkerRoot    Marker = "root"
	MarkerMention Marker = "mention"
)
