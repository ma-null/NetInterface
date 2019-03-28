package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net"
)


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

func GetNetInterfaceInfo(mng NetInterfaceManger, name string) (NetInterface, error) {
	netIf, err := mng.NetInterfaceByName(name)
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

func GetInterface(mng NetInterfaceManger) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
		fmt.Println("before nameIf")
		name := ps.ByName("name")
		if name == "" {
			http.Error(w, "name should be set", http.StatusBadRequest)
			return
		}

		fmt.Println("before GetInterfaceInfo")
		netIf, err := GetNetInterfaceInfo(mng, name)
		if err != nil {
			http.Error(w, "interface"+name+"was not found", http.StatusInternalServerError)
			return
		}

		fmt.Println("before json")
		w.Header().Set("Content-Type", "application/json")
		/*
		data, err := json.Marshal(netIf)
		if err != nil {
			fmt.Printf(err.Error())
		}*/
		if err = json.NewEncoder(w).Encode(netIf); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Printf(err.Error())
			return
		}
		fmt.Println("after GetInterfaceInfo/n")
	}
}