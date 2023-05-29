package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	TelegramBot TelegramBotCfg `yaml:"telegram_bot"`
	WhiteList   []int64        `yaml:"white_list"`
	GRPC        GRPC           `yaml:"grpc"`
}

type GRPC struct {
	Port string `yaml:"port"`
}

type TelegramBotCfg struct {
	Token             string `json:"token" yaml:"token"`
	CreatorUserID     int64  `json:"creator_user_id" yaml:"creator_user_id"`
	UpdatesTimeoutSec int    `json:"updates_timeout_sec" yaml:"updates_timeout_sec"`
}

func Load(configFile string, cfg interface{}) error {
	fileData, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("can't read config file: %w", err)
	}

	fileData = []byte(os.ExpandEnv(string(fileData)))

	if err = yaml.Unmarshal(fileData, cfg); err != nil {
		return fmt.Errorf("can't unmarshal config data: %w", err)
	}

	if v, ok := cfg.(interface{ Validate() error }); ok {
		if err = v.Validate(); err != nil {
			return fmt.Errorf("invalid config: %w", err)
		}
	}

	return nil
}
