package imports

import (
	"encoding/json"
	"fmt"
	"linkyard/internal/links"
	"net/http"

	"github.com/gofrs/uuid"
)

type Handler struct {
	linksStore *links.Store
	uuidGen    *uuid.Gen
}

func NewHandler(linksStore *links.Store, uuidGen *uuid.Gen) *Handler {
	return &Handler{
		linksStore: linksStore,
		uuidGen:    uuidGen,
	}
}

func (h *Handler) HandleImportLink(w http.ResponseWriter, r *http.Request) {
	var rawLinks FirefoxLink
	err := json.NewDecoder(r.Body).Decode(&rawLinks)
	if err != nil {
		fmt.Println("failed to decode:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var newLinks []links.Link
	rawLinks.goThroughLinks(nil, &newLinks, h.uuidGen)

	jsonResult, err := json.Marshal(newLinks)
	if err != nil {
		fmt.Println("failed to marshal result:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(jsonResult))

}
