package main

import (
	"flag"
	"log"

	"github.com/VladimirBlinov/MailSender/internal/app"
	"github.com/VladimirBlinov/MailSender/internal/config"
)

var (
	templatePath      string
	broadcastListPath string
	delay             int
	configPath        string
	subject           string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/", "path to config file")
	flag.StringVar(&subject, "subject", "new subject", "broadcast subject")
	flag.StringVar(&templatePath, "template-path", "template/template.html", "path to template file")
	flag.StringVar(&broadcastListPath, "broadcastList-path", "broadcastList.txt", "path to broadcastList file")
	flag.IntVar(&delay, "delay", 0, "delay before send")
}

func main() {
	flag.Parse()

	cfg, err := config.GetConf(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	cfg.Subject = subject
	cfg.TemplatePath = templatePath
	cfg.Delay = delay
	cfg.BroadcastListPath = broadcastListPath

	app.Run(cfg)
}
