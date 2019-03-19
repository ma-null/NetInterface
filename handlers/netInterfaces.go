package handlers

import (
	"encoding/json"
    "github.com/julienschmidt/httprouter"
	"net/http"
	"net"
)

type ifResponse struct {
	AllIntr[] string `json:"Interfaces"`
}


func netInterfaceNames() (ifResponse, error) {
	myInterfaces, err := net.Interfaces()
	if err != nil {
		return ifResponse{}, err
	}

	ifNames := make([]string, len(myInterfaces)) 
	for i := range myInterfaces {
		ifNames[i] = myInterfaces[i].Name
	}
	return ifResponse{ifNames}, nil
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