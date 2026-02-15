package links

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{
		store: store,
	}
}

// by adding (h *Handler) we're saying this function belongs to the Handler struct
func (h *Handler) HandleGetLinks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	links, err := h.store.getLinks(ctx)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(links)
}
