package store

import "github.com/VladimirBlinov/MailSender/internal/model"

type SubscriberRepo interface {
	GetSubscribers() ([]*model.Subscriber, error)
}
