package win

import (
	"context"
	"github.com/99-66/NaiaBackendApi/repositories"
	"os"
)

type StopWord struct {
	Word string `json:"word"`
} //@name StopWord

// Set 불용어를 등록한다
func (w *StopWord) Set() error {
	return setStopWords(w.Word)
}

// setStopWords Redis에 불용어를 등록한다
func setStopWords(word string) error {
	result := repositories.Connections.REDIS.SAdd(context.TODO(), os.Getenv("REDIS_KEY"), word)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
