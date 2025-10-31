package test

import (
    "fmt"
    "bytes"
    "encoding/json"
    "net/http"
	"testing"
	"log"
)

func TestWithDatabase(t *testing.T) {
    tests := []struct {
		url          string
		method       string
		params       map[string]interface{}
		responseCode int
	}{
		{"/region", http.MethodPost, map[string]interface{}{
            "name":"московская область",
            "slug":"moskovskaya-oblast",
            "number":1,
        }, http.StatusCreated},
        {"/region/1", http.MethodGet, map[string]interface{}{}, http.StatusOK},
        {"/regions", http.MethodGet, map[string]interface{}{}, http.StatusOK},
        {"/region/1", http.MethodPut, map[string]interface{}{
            "name":"московская область",
            "slug":"moskovskaya-oblast",
            "number":99,
        }, http.StatusOK},
		{"/city", http.MethodPost, map[string]interface{}{
		    "name":"Москва",
		    "slug":"msk",
		    "region_id":1,
		    "is_capital":true,
		    "type":"Село",
		    "latitude":55.75,
		    "longitude":37.61,
		    "time_zone":"Europe/Moscow",
		    "population":132,
		}, http.StatusCreated},
		{"/city/1", http.MethodGet, map[string]interface{}{}, http.StatusOK},
		{"/cities", http.MethodGet, map[string]interface{}{}, http.StatusOK},
		{"/city/1", http.MethodPut, map[string]interface{}{
            "name":"москва",
            "slug":"msk",
            "region_id":1,
            "is_capital":true,
            "type":"Город",
            "latitude":55.751244,
            "longitude":37.618423,
            "time_zone":"Europe/Moscow",
            "population":13274285,
        }, http.StatusOK},
		{"/city/1", http.MethodDelete, map[string]interface{}{}, http.StatusNoContent},
		{"/region/1", http.MethodDelete, map[string]interface{}{}, http.StatusNoContent},
		{"/sync", http.MethodPost, map[string]interface{}{}, http.StatusOK},
	}

    for _, test := range tests {
        jsonData, err := json.Marshal(test.params)
        fmt.Println(string(jsonData))
        body := bytes.NewBuffer(jsonData)
        req, err := http.NewRequest(test.method, "http://localhost:8080" + test.url, body)
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

        if resp.StatusCode != test.responseCode {
            t.Errorf("URL %s error: actual %d, expected %d", test.url, resp.StatusCode, test.responseCode)
        }
	}
}
