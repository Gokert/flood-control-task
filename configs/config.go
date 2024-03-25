package configs

import (
	"github.com/spf13/viper"
)

type DbRedisCfg struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DbNumber int    `yaml:"db"`
	Timer    int    `yaml:"timer"`
}

type DbFloodCfg struct {
	TimeDif           float64 `yaml:"timeDif"`
	MaxRequestCurrent int     `yaml:"maxRequestCurrent"`
	CountRequest      int     `yaml:"countRequest"`
	TimeSleep         float64 `yaml:"timeSleep"`
}

func GetRedisConfig() (*DbRedisCfg, error) {
	v := viper.GetViper()
	v.AutomaticEnv()

	cfg := &DbRedisCfg{
		Host:     v.GetString("REDIS_ADDR"),
		Password: v.GetString("REDIS_PASSWORD"),
		DbNumber: v.GetInt("REDIS_DB"),
		Timer:    v.GetInt("REDIS_TIMER"),
	}

	return cfg, nil
}

func GetFloodConfig() (*DbFloodCfg, error) {
	v := viper.GetViper()
	v.AutomaticEnv()

	cfg := &DbFloodCfg{
		TimeDif:           v.GetFloat64("FLOOD_TIME_DIF_MIL"),
		MaxRequestCurrent: v.GetInt("FLOOD_MAX_REQUEST"),
		CountRequest:      v.GetInt("FLOOD_COUNT_REQUEST"),
		TimeSleep:         v.GetFloat64("FLOOD_TIME_SLEEP_MIL"),
	}

	return cfg, nil
}
