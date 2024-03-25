package main

import (
	"task/configs"
	"task/configs/logger"
	"task/listen"
	"task/usecase"
)

func main() {
	log := logger.GetLogger()

	redisCfg, err := configs.GetRedisConfig("configs/redis_db.yaml")
	if err != nil {
		log.Error("Create redis config error: ", err)
		return
	}

	floodCfg, err := configs.GetFloodConfig("configs/flood_config.yaml")
	if err != nil {
		log.Error("Create flood config error: ", err)
		return
	}

	core, err := usecase.GetCore(floodCfg, redisCfg, log)
	if err != nil {
		log.Error("Create core error: ", err)
		return
	}

	log.Info("App running...")
	err = listen.ListenAndServ(core, floodCfg, log)
	if err != nil {
		log.Error("Listen error: ", err)
		return
	}
}
