package handlers

import (
	"encoding/json"
    "github.com/julienschmidt/httprouter"
	"net/http"
)

const ApiVersion = "v1"

type verResponse struct {
	Version string `json:"version"`
}

func GetVersion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ver := verResponse{ApiVersion}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


