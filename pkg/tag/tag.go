package tag

type TagType string

type Tag interface {
	Encode() []byte
}

const (
	TagTypeEvent  TagType = "e"
	TagTypePubKey TagType = "p"
)

type EventTag struct {
	Type     string
	EventID  []byte
	RelayURL []byte
}

func (t *EventTag) Encode() []byte {
	data := make([]byte, 0)
	data = append(data, []byte(TagTypeEvent)...)
	data = append(data, t.EventID...)
	data = append(data, t.RelayURL...)

	return data
}

type PubKeyTag struct {
	Type     TagType
	PubKey   []byte
	RelayURL []byte
}

func (t *PubKeyTag) Encode() []byte {
	data := make([]byte, 0)
	data = append(data, []byte(TagTypePubKey)...)
	data = append(data, t.PubKey...)
	data = append(data, t.RelayURL...)

	return data
}
