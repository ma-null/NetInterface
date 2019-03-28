package handlers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net"
	"testing"
)

func  mockNetInterfaceByName(name string) (*net.Interface, error) {
	return &net.Interface{
		1,
		1500,
		"en0",
		[]byte{2},
		1}, nil
}

func  mockAddrs(netIf net.Interface) ([]net.Addr, error) {
	return []net.Addr{&net.IPAddr{[]byte{2}, "192.0.2.1:25"}}, nil
}

func TestGetInterface(t *testing.T) {
	mng := NewmockNetIfManager(mockNetInterfaceByName, mockNetInterfaces, mockAddrs)
	rd, wr := io.Pipe()
	defer rd.Close()

	w := MyResponseWriter{wr}
	ps := httprouter.Params{{"name", "en0"}}

	go func() {
		GetInterface(mng)(&w ,nil, ps)
		if err := w.Writer.Close(); err != nil {
			t.Error(err.Error())
		}
	}()

	respNetIf := NetInterface{}
	body, err := ioutil.ReadAll(rd)
	if err != nil {
		t.Error(err.Error())
	}

	if err := json.Unmarshal(body, &respNetIf); err != nil {
		t.Error(err.Error())
	}

	if respNetIf.Name != "en0" {
		t.Error("Expected: " + "en0" + ", got: " + respNetIf.Name)
	}
}
