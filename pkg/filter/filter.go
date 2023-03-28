package filter

func New(opt *FilterOptions) *Filter {
	return &Filter{}
}

type FilterOptions struct {
	IDs        []string `json:"ids"`
	Authors    []string `json:"authors"`
	Kinds      []string `json:"kinds"`
	EventIDs   []string `json:"#e"`
	PublicKeys []string `json:"#p"`
	Since      string   `json:"since"`
	Until      string   `json:"until"`
	Limit      uint     `json:"limit"`
}

type Filter struct {
	IDs        []string `json:"ids"`
	Authors    []string `json:"authors"`
	Kinds      []string `json:"kinds"`
	EventIDs   []string `json:"#e"`
	PublicKeys []string `json:"#p"`
	Since      string   `json:"since"`
	Until      string   `json:"until"`
	Limit      uint     `json:"limit"`
}
