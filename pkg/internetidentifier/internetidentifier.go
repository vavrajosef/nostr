package internetidentifier

func New() *InternetIdentifier {
	return &InternetIdentifier{}
}

type InternetIdentifier struct {
	Names  map[string]string   `json:"names,omitempty"`
	Relays map[string][]string `json:"relays,omitempty"`
}
