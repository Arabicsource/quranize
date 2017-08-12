package api

import (
	"encoding/json"
	"net/http"

	"github.com/alpancs/quranize/service"
)

func Encode(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	json.NewEncoder(w).Encode(service.Encode(keyword))
}