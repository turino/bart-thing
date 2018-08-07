package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

// GetWebRequest is the interface
type GetWebRequest interface {
	FetchBytes(url string) ([]byte, error)
}

// LiveGetWebRequest isn't testWebRequest
type LiveGetWebRequest struct {
}

// FetchBytes returns bytes for given url
func (LiveGetWebRequest) FetchBytes(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "bart-thing")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
