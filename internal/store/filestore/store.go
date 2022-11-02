package filestore

import "github.com/VladimirBlinov/MailSender/internal/store"

type Store struct {
	filePath       string
	subscriberRepo *SubscriberRepo
}

func NewStore(fp string) *Store {
	return &Store{
		filePath: fp,
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
