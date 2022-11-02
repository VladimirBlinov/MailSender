package service

import (
	"log"
	"sync"
	"time"

	"github.com/VladimirBlinov/MailSender/internal/store"
	"github.com/VladimirBlinov/MailSender/pkg/email"
)

type Service struct {
	Store  store.Store
	Sender email.Sender
}

func NewService(store store.Store, sender email.Sender) *Service {
	return &Service{
		Store:  store,
		Sender: sender,
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
				log.Printf("Error generate email to %s", sub.Email)
			}

			err = s.Sender.Send(sendInput)
			if err != nil {
				log.Printf("Error sent to email %s", sub.Email)
			}
		}(wg)

		wg.Wait()
	}

	return nil
}
