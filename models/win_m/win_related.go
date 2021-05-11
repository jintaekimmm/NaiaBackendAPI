package win_m

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type WRelatedWords struct {
	Code    int `json:"code"`
	Message struct {
		Nodes []RNodes           `json:"nodes"`
		Links []RLinks           `json:"links"`
		Rank  map[string]float64 `json:"rank"`
	} `json:"message"`
}

type RNodes struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Size float64 `json:"_size"`
}

type RLinks struct {
	Sid string `json:"sid"`
	Tid string `json:"tid"`
}

type WordsResponse struct {
	Nodes []RNodes           `json:"nodes"`
	Links []RLinks           `json:"links"`
	Rank  map[string]float64 `json:"rank"`
}

type WRelatedTweets struct {
	Code    int       `json:"code"`
	Message []RTweets `json:"message"`
}

type RTweets struct {
	CreatedAt time.Time `json:"createdAt"`
	Text      string    `json:"text"`
}

func (w *WRelatedWords) RelatedWords(word string) (WordsResponse, error) {
	baseUrl := os.Getenv("RELATED_API")
	if baseUrl == "" {
		return WordsResponse{}, fmt.Errorf("url is empty")
	}
	url := fmt.Sprintf("%s/related/%s", baseUrl, word)

	res, err := newHttpRequest("GET", url)
	if err != nil {
		return WordsResponse{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(w)
	if err != nil {
		return WordsResponse{}, fmt.Errorf("json decode error")
	}

	return WordsResponse{
		Nodes: w.Message.Nodes,
		Links: w.Message.Links,
		Rank:  w.Message.Rank,
	}, nil
}

func (w *WRelatedTweets) RelatedTweets(word string) ([]RTweets, error) {
	baseUrl := os.Getenv("RELATED_API")
	if baseUrl == "" {
		return []RTweets{}, fmt.Errorf("url is empty")
	}
	url := fmt.Sprintf("%s/related/list/%s", baseUrl, word)

	res, err := newHttpRequest("GET", url)
	if err != nil {
		return []RTweets{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(w)
	if err != nil {
		return []RTweets{}, fmt.Errorf("json decode error")
	}

	return w.Message, nil
}

func newHttpRequest(method, url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("constructor error")
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request error")
	}

	return res, nil
}
