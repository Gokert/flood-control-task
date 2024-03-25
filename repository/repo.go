package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"task/configs"
	"time"
)

type Repo struct {
	db *redis.Client
}

func GetRepo(cfg *configs.DbRedisCfg, log *logrus.Logger) (*Repo, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.DbNumber,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Error("Ping redis error: ", err)
		return nil, err
	}

	log.Info("Redis created successful")
	return &Repo{db: redisClient}, nil
}

func (r *Repo) Set(ctx context.Context, key, value string, exp time.Duration) error {
	messageKey := fmt.Sprintf("%s:message:%d", key, time.Now().UnixNano()/int64(time.Millisecond))

	err := r.db.Set(ctx, messageKey, value, exp).Err()
	if err != nil {
		return fmt.Errorf("set db error: %s", err.Error())
	}

	_, err = r.Get(ctx, key)
	if err != nil {
		return fmt.Errorf("not found value: %s", err.Error())
	}

	return nil
}

func (r *Repo) Get(ctx context.Context, key string) ([]string, error) {
	keys, err := r.db.Keys(ctx, fmt.Sprintf("%s:message:*", key)).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, fmt.Errorf("get db error: %s ", err.Error())
		}
	}

	return keys, nil
}
