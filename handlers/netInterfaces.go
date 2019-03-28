package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type IfResponse struct {
	AllIntr[] string `json:"Interfaces"`
}

func netInterfaceNames(mng NetInterfaceManger) (IfResponse, error) {
	netIf, err := mng.NetInterfaces()
	if err != nil {
		return IfResponse{}, err
	}
	ifNames := make([]string, len(netIf)) 
	for i := range netIf {
		ifNames[i] = netIf[i].Name
	}
	return IfResponse{ifNames}, nil
}

func GetInterfaces(mng NetInterfaceManger) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		netIf, err := netInterfaceNames(mng)

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
}