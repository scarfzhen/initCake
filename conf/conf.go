package conf

import (
	"fmt"
	"initCake/pkg/cfg"
	"initCake/pkg/httpx"
	"initCake/pkg/logx"
	"initCake/pkg/ormx"
)

type ConfigType struct {
	Global GlobalConfig
	Log    logx.Config
	HTTP   httpx.Config
	DB     ormx.DBConfig
}

type GlobalConfig struct {
	RunMode string
}

func InitConfig(configDir, cryptoKey string) (*ConfigType, error) {
	var config = new(ConfigType)
	if err := cfg.LoadConfigByDir(configDir, config); err != nil {
		return nil, fmt.Errorf("failed to load configs of directory: %s error: %s", configDir, err)
	}
	err := decryptConfig(config, cryptoKey)
	if err != nil {
		return nil, err
	}
	return config, nil

}
