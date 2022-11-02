package smtpsender

import (
	"fmt"

	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/VladimirBlinov/MailSender/pkg/email"
)

type SMTPSender struct {
	from   string
	pass   string
	host   string
	port   int
	logger *logrus.Logger
}

func NewSMTPSender(from, pass, host string, port int) (*SMTPSender, error) {
	logger := logrus.New()
	if !email.IsEmailValid(from) {
		return nil, errors.New(fmt.Sprintf("invalid from email: %s", from))
	}

	return &SMTPSender{from: from, pass: pass, host: host, port: port, logger: logger}, nil
}

func (s *SMTPSender) Send(input email.SendEmailInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetBody("text/html", input.Body)

	s.logger.Infof("Sender - host: %s, port: %v", s.host, s.port)

	dialer := gomail.NewDialer(s.host, s.port, s.from, s.pass)
	if err := dialer.DialAndSend(msg); err != nil {
		return errors.Wrap(err, "failed to sent email via smtp")
	}

	return nil
}
