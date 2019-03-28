package handlers

import (
	"encoding/json"
    "github.com/julienschmidt/httprouter"
	"net/http"
)

const ApiVersion = "v1"

type VerResponse struct {
	Version string `json:"version"`
}

func GetVersion(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ver := VerResponse{ApiVersion}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ver); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


