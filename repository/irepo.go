package repository

import (
	"context"
	"time"
)

type IRepo interface {
	Set(ctx context.Context, key, value string, exp time.Duration) error
	Get(ctx context.Context, key string) ([]string, error)
}
