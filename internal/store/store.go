package store

type Store interface {
	Subscriber() SubscriberRepo
}
