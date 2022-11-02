package config

import (
	"github.com/spf13/viper"
)

type SMTPConfig struct {
	From string `yaml:"from"`
	Pass string `yaml:"pass"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	SMTP              SMTPConfig
	Subject           string
	TemplatePath      string
	BroadcastListPath string
	Delay             int
}

func GetConf(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("emailservice")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
