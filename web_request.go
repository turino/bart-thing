package bart

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// GetWebRequest is the interface
type GetWebRequest interface {
	FetchBytes(url string) []byte
}

// LiveGetWebRequest isn't testWebRequest
type LiveGetWebRequest struct {
}

// FetchBytes returns bytes for given url
func (LiveGetWebRequest) FetchBytes(url string) []byte {
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "bart-thing")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
