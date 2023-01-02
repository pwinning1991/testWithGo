package signal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fatalf("Handler() status = %d; want %d", resp.StatusCode, 200)
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q; want %q", contentType, "application/json")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(resp.Body) err = %s", err)
	}
	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal(resp.Body) err = %s", err)
	}
	if p.Age != 21 {
		t.Errorf("person.Age = %d; want %d", p.Age, 21)
	}
	if p.Name != "Phil" {
		t.Errorf("person.Name = %s; want %s", p.Name, "Phil")
	}
}
