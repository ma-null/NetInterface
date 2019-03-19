package handlers

import (
	"encoding/json"
    "github.com/julienschmidt/httprouter"
	"net/http"
	"net"
)

type IfResponse struct {
	AllIntr[] string `json:"Interfaces"`
}

func netInterfaceNames() (IfResponse, error) {
	netIf, err := net.Interfaces()
	if err != nil {
		return IfResponse{}, err
	}
	ifNames := make([]string, len(netIf)) 
	for i := range netIf {
		ifNames[i] = netIf[i].Name
	}
	return IfResponse{ifNames}, nil
}

func GetIntefaces(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	netIf, err := netInterfaceNames()	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(netIf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}