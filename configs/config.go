package configs

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type DbRedisCfg struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DbNumber int    `yaml:"db"`
	Timer    int    `yaml:"timer"`
}

type DbFloodCfg struct {
	TimeDif           int   `yaml:"timeDif"`
	MaxRequestCurrent int   `yaml:"maxRequestCurrent"`
	CountRequest      int   `yaml:"countRequest"`
	UserId            int64 `yaml:"userId"`
	TimeSleep         int   `yaml:"timeSleep"`
}

func GetRedisConfig(cfgPath string) (*DbRedisCfg, error) {
	v := viper.GetViper()
	v.SetConfigFile(cfgPath)
	v.SetConfigType(strings.TrimPrefix(filepath.Ext(cfgPath), "."))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &DbRedisCfg{
		Host:     v.GetString("host"),
		Password: v.GetString("password"),
		DbNumber: v.GetInt("db"),
		Timer:    v.GetInt("timer"),
	}

	return cfg, nil
}

func GetFloodConfig(cfgPath string) (*DbFloodCfg, error) {
	v := viper.GetViper()
	v.SetConfigFile(cfgPath)
	v.SetConfigType(strings.TrimPrefix(filepath.Ext(cfgPath), "."))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &DbFloodCfg{
		TimeDif:           v.GetInt("timeDif"),
		MaxRequestCurrent: v.GetInt("maxRequestCurrent"),
		CountRequest:      v.GetInt("countRequest"),
		UserId:            v.GetInt64("userId"),
		TimeSleep:         v.GetInt("timeSleep"),
	}

	return cfg, nil
}
