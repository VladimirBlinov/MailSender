package filestore

import (
	"github.com/VladimirBlinov/MailSender/internal/store"
	"github.com/sirupsen/logrus"
)

type Store struct {
	filePath       string
	subscriberRepo *SubscriberRepo
	logger         *logrus.Logger
}

func NewStore(fp string, logger *logrus.Logger) *Store {
	return &Store{
		filePath: fp,
		logger:   logger,
	}
}

func (s *Store) Subscriber() store.SubscriberRepo {
	if s.subscriberRepo != nil {
		return s.subscriberRepo
	}

	s.subscriberRepo = &SubscriberRepo{
		store: s,
	}
	return s.subscriberRepo
}
