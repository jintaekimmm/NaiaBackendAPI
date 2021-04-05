package models

type WinList struct {
	Word string `json:"word"`
	Count int64 `json:"count"`
}

type WinTag struct {
	Tag string `json:"tag"`
	Percent float64 `json:"percent"`
}