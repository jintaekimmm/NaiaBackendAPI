package services

import (
	"context"
	"github.com/99-66/NaiaBackendApi/repositories"
	"github.com/go-redis/redis/v8"
	"os"
	"time"
)

// StopWords 불용어 목록을 반환한다
func StopWords() ([]string, error) {
	return getStopWords(context.Background(), repositories.Connections.REDIS, os.Getenv("REDIS_KEY"))
}

// getStopWords Redis에서 불용어 목록을 읽어온다
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
