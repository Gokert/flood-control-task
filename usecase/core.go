package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"strconv"
	"task/configs"
	"task/repository"
	"time"
)

type FloodControl interface {
	Check(ctx context.Context, userID int64) (bool, error)
}

type Core struct {
	log  *logrus.Logger
	repo repository.IRepo
	cfg  *configs.DbFloodCfg
}

func GetCore(floodCfg *configs.DbFloodCfg, redisCfg *configs.DbRedisCfg, log *logrus.Logger) (*Core, error) {
	repo, err := repository.GetRepo(redisCfg, log)
	if err != nil {
		log.Errorf("get core error: %s", err.Error())
		return nil, err
	}

	core := &Core{
		log:  log,
		repo: repo,
		cfg:  floodCfg,
	}

	return core, nil
}

func (c *Core) Check(ctx context.Context, userID int64) (bool, error) {
	array, err := c.repo.Get(ctx, strconv.FormatInt(userID, 10))
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return false, fmt.Errorf("check get error: %s", err.Error())
		}
	}

	if len(array) >= c.cfg.MaxRequestCurrent {
		return false, nil
	}

	err = c.repo.Set(ctx, strconv.FormatInt(userID, 10), "", time.Duration(c.cfg.TimeDif)*time.Second)
	if err != nil {
		return false, fmt.Errorf("check set error: %s", err.Error())
	}

	return true, nil
}
