package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net"
	"strings"
	"testing"
)

func  mockNetInterfaces() ([]net.Interface, error) {
	return []net.Interface{{
		1,
		1500,
		"en0",
		[]byte{2},
		1},
		{2,
		1500,
		"ln2",
		[]byte{1},
		1}}, nil
}

func TestGetInterfaces(t *testing.T) {
	mng := NewmockNetIfManager(mockNetInterfaceByName, mockNetInterfaces, mockAddrs)
	rd, wr := io.Pipe()
	defer rd.Close()

	w := MyResponseWriter{wr}
	go func() {
		GetInterfaces(mng)(&w ,nil, nil)
		if err := w.Writer.Close(); err != nil {
			t.Error(err.Error())
		}
	}()

	respNetIf := IfResponse{}
	body, err := ioutil.ReadAll(rd)
	if err != nil {
		t.Error(err.Error())
	}

	if err := json.Unmarshal(body, &respNetIf); err != nil {
		t.Error(err.Error())
	}

	if len(respNetIf.AllIntr) != 2 {
		t.Error("Expected: 2 interfaces" + ", got: " + strings.Join( respNetIf.AllIntr, ","))
	}
}
