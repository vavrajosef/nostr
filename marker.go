package nostr

type MarkerType string

const (
	MarkerTypeReply   MarkerType = "reply"
	MarkerTypeRoot    MarkerType = "root"
	MarkerTypeMention MarkerType = "mention"
)
