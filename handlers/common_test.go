package handlers

import (
	"io"
	"net"
	"net/http"
)

type mockNetIfManager struct {
	//netIf NetInterface
	mockNetInterfaceByName func(name string) (*net.Interface, error)
	mockNetInterfaces      func() ([]net.Interface, error)
	mockAddrs              func(netIf net.Interface) ([]net.Addr, error)
}

func NewmockNetIfManager(
	mockNetInterfaceByName func(name string) (*net.Interface, error),
	mockNetInterfaces func() ([]net.Interface, error),
	mockAddrs func(netIf net.Interface) ([]net.Addr, error)) NetInterfaceManger {
	return &mockNetIfManager{
		mockNetInterfaceByName,
		mockNetInterfaces,
		mockAddrs}
}

func (m *mockNetIfManager) NetInterfaceByName(name string) (*net.Interface, error) {
	return mockNetInterfaceByName(name)
}

func (m *mockNetIfManager) NetInterfaces() ([]net.Interface, error) {
	return mockNetInterfaces()
}

func (m *mockNetIfManager) Addrs(netIf net.Interface) ([]net.Addr, error) {
	return mockAddrs(netIf)
}

type MyResponseWriter struct {
	Writer io.WriteCloser
}

func (w *MyResponseWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}

func (w *MyResponseWriter) WriteHeader(status int) {}

func (w *MyResponseWriter) Header() http.Header {
	return http.Header{}

}