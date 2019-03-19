package handlers

import (
	"encoding/json"
    "github.com/julienschmidt/httprouter"
	"net/http"
	"net"
)

type NetInterface struct {
	Name      string           `json:"name"`
	Hw_addr   net.HardwareAddr `json:"hw_addr"` 
	Inet_addr []string       `json:"inet_addr"`
	MTU       int              `json:"MTU"`
}

func netInterfaceAddr(netIf net.Interface) ([]string, error) {
	addrs, err := netIf.Addrs()
	if err != nil {
		return []string{}, err
	} 

	addrIps := make([]string, 0, len(addrs))
	for _, addr := range addrs {
		addrIps = append(addrIps, addr.String())
	} 
	return addrIps, nil
}

func GetNetInterfaceInfo(name string) (NetInterface, error) {
	netIf, err := net.InterfaceByName(name)
	if err != nil {
		return NetInterface{}, err
	} 

	addr, err := netInterfaceAddr(*netIf)
	if err != nil {
		return NetInterface{}, err
	} 
	return NetInterface{ 
		netIf.Name,
		netIf.HardwareAddr,
		addr,		
		netIf.MTU,
	}, nil
}

func GetInteface(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if name == "" {
		http.Error(w, "name should be set", http.StatusBadRequest) 
		return	
	}

	netIf, err := GetNetInterfaceInfo(name)
	if err != nil {
		http.Error(w, "interface"+name+"was not found", http.StatusInternalServerError) 
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(netIf); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}