package services

import (
	"encoding/json"
	"fmt"
	"github.com/99-66/NaiaBackendApi/models/win"
	"net/http"
	"os"
)

func RelatedWords(word string) (win.WordsResponse, error) {
	baseUrl := os.Getenv("RELATED_API")
	if baseUrl == "" {
		return win.WordsResponse{}, fmt.Errorf("url is empty")
	}
	url := fmt.Sprintf("%s/related/%s", baseUrl, word)

	res, err := newHttpRequest("GET", url)
	if err != nil {
		return win.WordsResponse{}, err
	}
	defer res.Body.Close()

	var relatedWords win.RelatedWords
	err = json.NewDecoder(res.Body).Decode(&relatedWords)
	if err != nil {
		return win.WordsResponse{}, fmt.Errorf("json decode error")
	}

	return win.WordsResponse{
		Nodes: relatedWords.Message.Nodes,
		Links: relatedWords.Message.Links,
		Rank:  relatedWords.Message.Rank,
	}, nil
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

func RelatedTweets(word string) ([]win.RelatedTweet, error) {
	baseUrl := os.Getenv("RELATED_API")
	if baseUrl == "" {
		return []win.RelatedTweet{}, fmt.Errorf("url is empty")
	}
	url := fmt.Sprintf("%s/related/list/%s", baseUrl, word)

	res, err := newHttpRequest("GET", url)
	if err != nil {
		return []win.RelatedTweet{}, err
	}
	defer res.Body.Close()

	var rTweets win.RelatedTweets
	err = json.NewDecoder(res.Body).Decode(&rTweets)
	if err != nil {
		return []win.RelatedTweet{}, fmt.Errorf("json decode error")
	}

	return rTweets.Message, nil
}
