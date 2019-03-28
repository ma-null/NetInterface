package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestGetVersion(t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:8080/version", nil)
	w := httptest.NewRecorder()

	GetVersion(w, r, nil)
	ver := VerResponse{}
	if err := json.Unmarshal(w.Body.Bytes(), &ver); err != nil {
		t.Error(err.Error())
	}
	if ver.Version != "v1" {
		t.Error("Invalid version: "+ver.Version+", expected: "+"v1")
	}
}