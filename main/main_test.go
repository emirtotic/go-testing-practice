package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/double?v=2", nil)

	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()

	doubleHandler(rec, req)

	response := rec.Result()

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got: %v", response.Status)
	}

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Errorf("Could not read response : %v", response)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))

	if err != nil {
		t.Errorf("Expected an integer; got %v", d)
	}

	if d != 4 {
		t.Fatalf("Expected double to be 4 but it is %v", d)
	}

}
