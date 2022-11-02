package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Subscriber struct {
	Name     string
	LastName string
	BirthDay time.Time
	Email    string
}

func NewSubscriber(n string, ln string, bd time.Time, e string) *Subscriber {
	return &Subscriber{
		Name:     n,
		LastName: ln,
		BirthDay: bd,
		Email:    e,
	}
}

func (s *Subscriber) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required, validation.Length(2, 50)),
		validation.Field(&s.LastName, validation.Required, validation.Length(2, 50)),
		validation.Field(&s.BirthDay, validation.Required),
		validation.Field(&s.Email, validation.Required, is.Email),
	)
}
