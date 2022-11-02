package service

import (
	"sync"
	"time"

	"github.com/VladimirBlinov/MailSender/internal/store"
	"github.com/VladimirBlinov/MailSender/pkg/email"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Store  store.Store
	Sender email.Sender
	Logger *logrus.Logger
}

func NewService(store store.Store, sender email.Sender, logger *logrus.Logger) *Service {

	return &Service{
		Store:  store,
		Sender: sender,
		Logger: logger,
	}
}

func (s *Service) RunBroadcast(emailTempl string, subject string, delay int) error {
	subscribers, err := s.Store.Subscriber().GetSubscribers()
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(delay) * time.Second)
	wg := &sync.WaitGroup{}

	for _, sb := range subscribers {
		sub := sb

		wg.Add(1)
		go func(waiter *sync.WaitGroup) {
			defer waiter.Done()
			sendInput := email.SendEmailInput{Subject: subject, To: sub.Email}

			if err := sendInput.GenerateBodyFromHTML(emailTempl, sub); err != nil {
				s.Logger.Errorf("Service error: Error generate email to %s: %s", sub.Email, err)
			}

			err = s.Sender.Send(sendInput)
			if err != nil {
				s.Logger.Errorf("Service error: Error sent to email %s: %s", sub.Email, err)
			}
		}(wg)

		wg.Wait()
	}

	return nil
}
