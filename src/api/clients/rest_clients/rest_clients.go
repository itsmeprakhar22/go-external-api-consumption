package rest_clients

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Get(url string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodGet, url, bytes.NewReader(jsonBytes))
	request.Header = headers
	client := http.Client{}

	return client.Do(request)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers
	client := http.Client{}

	return client.Do(request)
}
