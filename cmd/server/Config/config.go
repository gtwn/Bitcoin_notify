package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug        	bool
	// Port         	int    `required:"true"`
	BitAPI			string `envconfig:"api_ticker" require:"true"`
	LineNoti		string `envconfig:"line_noti" require:"true"`
	ChannelAccessToken		string `envconfig:"channel_access_token" require:"true"`
}

func Read() (Config, error) {
	var cfg Config
	err := envconfig.Process("BIT", &cfg)
	return cfg, err
}
