package rest_clients

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Get(url string, body interface{}, headers http.Header) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = headers
	request.Header.Add("Content-Type", `application/json;charset=utf-8`)
	client := http.DefaultClient

	return client.Do(request)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers
	request.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	request.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	client := http.Client{}

	return client.Do(request)
}
