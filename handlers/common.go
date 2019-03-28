package handlers

import (
	"net"
)

type NetInterface struct {
	Name      string           `json:"name"`
	Hw_addr   net.HardwareAddr `json:"hw_addr"`
	Inet_addr []string       `json:"inet_addr"`
	MTU       int              `json:"MTU"`
}

type (
	NetInterfaceManger interface {
		NetInterfaceByName(name string) (*net.Interface, error)
		NetInterfaces()([]net.Interface, error)
		Addrs(netIf net.Interface) ([]net.Addr, error)
	}

	netInterfaceManger struct {}
)

func NewNetInterfaceManger() NetInterfaceManger {
	return &netInterfaceManger{}
}

func (m *netInterfaceManger) NetInterfaces()([]net.Interface, error)  {
	return net.Interfaces()
}
func (m *netInterfaceManger) Addrs(netIf net.Interface) ([]net.Addr, error) {
	return netIf.Addrs()
}
func (m *netInterfaceManger) NetInterfaceByName(name string) (*net.Interface, error) {
	return net.InterfaceByName(name)
}
