package win

type Tag struct {
	Tag     string  `json:"tag"`
	Percent float64 `json:"percent"`
} //@name Tag

type TagCount struct {
	Tag   string `json:"tag"`
	Count int64  `json:"count"`
} //@name TagCount
