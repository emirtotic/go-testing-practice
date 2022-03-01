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

	var tt = []struct {
		name   string
		value  string
		double int
		status int
		error  string
	}{
		{name: "Double of 2", value: "2", status: http.StatusOK, double: 4},
		{name: "Missing value", value: "", status: http.StatusBadRequest, error: "Missing value"},
	}
	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "localhost:8080/double?v="+tc.value, nil)

			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			rec := httptest.NewRecorder()
			doubleHandler(rec, req)
			response := rec.Result()
			defer response.Body.Close()

			if response.StatusCode != tc.status {
				t.Errorf("Expected status %v; got: %v", tc.status, response.Status)
			}

			b, err := ioutil.ReadAll(response.Body)

			if err != nil {
				t.Fatalf("Could not read response : %v", response)
			}

			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))

			if tc.error != "" {
				if response.StatusCode != http.StatusBadRequest {
					t.Errorf("Expected Status Bad Request; got %v", response.StatusCode)
				}

				if msg := string(bytes.TrimSpace(b)); msg != tc.error {
					t.Errorf("Expected message %q; got %q", tc.error, msg)
				}
				return
			}

			if d != 4 {
				t.Fatalf("Expected double to be 4 but it is %v", d)
			}

		})

	}
}
