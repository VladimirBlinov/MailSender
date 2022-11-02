package config

import (
	"github.com/spf13/viper"
)

type SMTPConfig struct {
	From string
	Pass string 
	Host string 
	Port int    
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

	conf := &Config{
		&SMTPConfig{
			From: viper.GetString("smtp.from"),
			Pass: viper.GetString("smtp.pass"),
			Pass: viper.GetString("smtp.pass"),
		}
	}

	// err = viper.Unmarshal(conf)
	// if err != nil {
	// 	return nil, err
	// }

	return conf, nil
}
