package listen

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"task/configs"
	"task/usecase"
	"time"
)

func ListenAndServ(c usecase.FloodControl, cfg *configs.DbFloodCfg, log *logrus.Logger) error {
	userId := rand.Int63n(1000)

	for i := 0; i < cfg.CountRequest; i++ {
		log.Infof("Request number: %d", i)

		result, err := c.Check(context.Background(), userId)
		if err != nil {
			return fmt.Errorf("listen check error: %s", err.Error())
		}

		if !result {
			return fmt.Errorf("result error: FLOOD")
		}

		time.Sleep(time.Duration(cfg.TimeSleep) * time.Millisecond)
	}

	return nil
}
