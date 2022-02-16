package redis

import (
	"fmt"
	"strconv"

	"github.com/andrei-dascalu/shortener/src/shortener"
	"github.com/go-redis/redis"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	opts, err := redis.ParseURL(redisURL)

	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	_, err = client.Ping().Result()

	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewRedisRepository(redisURL string) (shortener.LinkRepository, error) {
	repo := &redisRepository{}

	client, err := newRedisClient(redisURL)

	if err != nil {
		return nil, err
	}

	repo.client = client

	return repo, nil
}

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (r *redisRepository) Find(code string) (*shortener.LinkRedirect, error) {
	redirect := &shortener.LinkRedirect{}

	key := r.generateKey(code)

	data, err := r.client.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}

	createdAt, err := strconv.ParseInt(data["created_at"], 10, 64)

	if err != nil {
		return nil, err
	}

	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = createdAt

	return redirect, nil
}

func (r *redisRepository) Store(link *shortener.LinkRedirect) error {
	key := r.generateKey(link.Code)

	data := map[string]interface{}{
		"code":       link.Code,
		"url":        link.URL,
		"created_at": link.CreatedAt,
	}

	r.client.HMSet(key, data).Result()

	return nil
}
