package app

import (
	"github.com/VladimirBlinov/MailSender/internal/config"
	"github.com/VladimirBlinov/MailSender/internal/service"
	"github.com/VladimirBlinov/MailSender/internal/store/filestore"
	"github.com/VladimirBlinov/MailSender/pkg/email/smtpsender"
	"github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) {
	logger := logrus.New()

	emailSender, err := smtpsender.NewSMTPSender(cfg.SMTP.From, cfg.SMTP.Pass, cfg.SMTP.Host, cfg.SMTP.Port)
	if err != nil {
		logger.Error(err)
		return
	}

	store := filestore.NewStore(cfg.BroadcastListPath, logger)
	srvc := service.NewService(store, emailSender, logger)
	srvc.RunBroadcast(cfg.TemplatePath, cfg.Subject, cfg.Delay)
}
