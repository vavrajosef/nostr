package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-nostr/go-nostr/internal/core"
)

func NewGetHealthHandler(s *core.Service) *GetHealthHandler {
	return &GetHealthHandler{s}
}

type GetHealthHandler struct {
	s *core.Service
}

func (h *GetHealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &core.GetHealthRequest{}
	res, _ := h.s.GetHealth(req)
	byt, _ := json.Marshal(res)

	// NOTE: TBD
	w.Header().Add("Content-Type", "application/json")

	// NOTE: TBD
	w.Write(byt)
}
