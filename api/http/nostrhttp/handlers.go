package nostrhttp

import (
	"encoding/json"
	"net/http"

	"github.com/go-nostr/go-nostr/internal/core"
)

func newGetHealthHandler(svc *core.Service) *getHealthHandler {
	return &getHealthHandler{svc}
}

type getHealthHandler struct {
	svc *core.Service
}

func (h *getHealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &core.GetHealthRequest{}
	res, _ := h.svc.GetHealth(req)
	byt, _ := json.Marshal(res)

	// NOTE: TBD
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	// NOTE: TBD
	w.Write(byt)
}

func newGetInternetIdentifier(svc *core.Service) *getInternetIdentifier {
	return &getInternetIdentifier{svc}
}

type getInternetIdentifier struct {
	svc *core.Service
}

func (h *getInternetIdentifier) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, _ := h.svc.GetInternetIdentifier(&core.GetInternetIdentifierRequest{})
	byt, _ := json.Marshal(res)

	// NOTE: TBD
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	// NOTE: TBD
	w.Write(byt)
}
