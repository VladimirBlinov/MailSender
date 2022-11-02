package config

import (
	"github.com/spf13/viper"
)

type SMTPConfig struct {
	From string `mapstructure:"from"`
	Pass string `mapstructure:"pass"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
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

	smtpconf := &SMTPConfig{
		From: viper.GetString("smtp.from"),
		Pass: viper.GetString("smtp.pass"),
		Host: viper.GetString("smtp.host"),
		Port: viper.GetInt("smtp.port"),
	}

	conf := &Config{
		SMTP: *smtpconf,
	}

	return conf, nil
}
