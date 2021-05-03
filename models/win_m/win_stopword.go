package win_m

import (
	"context"
	"github.com/99-66/NaiaBackendApi/repositories"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

type WStopWord struct {
	Word string `json:"word"`
}

// List 단어 집계시에 제외할 불용어 목록을 읽어온다
// TODO : REDIS_KEY 변수화
func (w *WStopWord) List() ([]string, error) {
	return getStopWords(context.Background(), repositories.Connections.REDIS, os.Getenv("REDIS_KEY"))
}

func getStopWords(ctx context.Context, redis *redis.Client, key string) ([]string, error) {
	// 컨텍스트 타임아웃 설정
	localCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	result := redis.SMembers(localCtx, key)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return result.Val(), nil
}

// Set 불용어를 등록한다
// TODO : REDIS_KEY 변수화
func (w *WStopWord) Set() error {
	return setStopWords(w.Word)
}

func setStopWords(word string) error {
	result := repositories.Connections.REDIS.SAdd(context.TODO(), os.Getenv("REDIS_KEY"), word)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
