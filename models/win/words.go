package win

type WordCount struct {
	Date      []string `json:"date"`
	Total     []int64  `json:"total"`
	SNS       []int64  `json:"sns"`
	Article   []int64  `json:"article"`
	Community []int64  `json:"community"`
} //@name WordCount

type WordCloud struct {
	Words []interface{} `json:"words"`
} //@name WordCloud
