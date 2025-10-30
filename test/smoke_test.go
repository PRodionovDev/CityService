package test

import (
    "net/http"
	"testing"
	"log"
)

func TestRequests(t *testing.T) {
    tests := []struct {
		url    string
		method string
	}{
		{"/", http.MethodGet},
		{"/cities", http.MethodGet},
		{"/regions", http.MethodGet},
		{"/city/1", http.MethodGet},
		{"/region/1", http.MethodGet},
	}

    for _, tt := range tests {
        req, err := http.NewRequest(tt.method, "http://localhost:8080" + tt.url, nil)
        if err != nil {
        	log.Fatal(err)
        	return
        }

        req.Header.Add("Authorization", "test")
        req.Header.Add("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
        	log.Fatal(err)
        	return
        }

        expected := http.StatusOK
        actual := resp.StatusCode

        if actual != expected {
            t.Errorf("URL %s error: actual %d, expected %d", tt.url, actual, expected)
        }
	}
}
