package api

import (
	"autosuggest/internal/app"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	App *app.Service
}

func NewHandler(app *app.Service) *Handler {
	return &Handler{App: app}
}

func (h *Handler) Suggest(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	kStr := r.URL.Query().Get("k")
	k, _ := strconv.Atoi(kStr)
	if k == 0 {
		k = 5
	}
	results := h.App.Suggest(q, k)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func (h *Handler) AddPhrase(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
		Freq int64  `json:"freq"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	h.App.AddOrUpdate(req.Text, req.Freq)
	w.WriteHeader(http.StatusOK)
}
