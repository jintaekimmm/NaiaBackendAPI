package win

import "time"

type RelatedWords struct {
	Code    int `json:"code"`
	Message struct {
		Nodes []RNodes           `json:"nodes"`
		Links []RLinks           `json:"links"`
		Rank  map[string]float64 `json:"rank"`
	} `json:"message"`
} //@name RelatedWords

type RNodes struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Size float64 `json:"_size"`
} //@name RelatedNodes

type RLinks struct {
	Sid string `json:"sid"`
	Tid string `json:"tid"`
} //@name RelatedLinks

type WordsResponse struct {
	Nodes []RNodes           `json:"nodes"`
	Links []RLinks           `json:"links"`
	Rank  map[string]float64 `json:"rank"`
} //@name WordsResponse

type RelatedTweets struct {
	Code    int            `json:"code"`
	Message []RelatedTweet `json:"message"`
} //@name RelatedTweets

type RelatedTweet struct {
	CreatedAt time.Time `json:"createdAt"`
	Text      string    `json:"text"`
} //@name RelatedTweet
