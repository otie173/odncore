package handler

import (
	"encoding/json"
	"net/http"

	"github.com/otie173/odncore/api/setup"
)

func respondJSON(res http.ResponseWriter, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(data); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(api setup.API) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		info := api.GetServer().GetInfo()
		respondJSON(res, info)
	}
}
