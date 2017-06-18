package route

import (
	"encoding/json"
	"net/http"

	"github.com/alpancs/quranize/service"
	"github.com/pressly/chi"
)

func Encode(w http.ResponseWriter, r *http.Request) {
	text := chi.URLParam(r, "text")
	encodeds := service.Encode(text)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(encodeds)
}
