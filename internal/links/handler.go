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

func (h *Handler) HandleCreateLink(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var link Link
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		fmt.Printf("decode error: %v", err)
		return
	}

	newLink, err := h.store.createLink(ctx, link)
	if err != nil {
		fmt.Printf("insert error: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newLink)
}
