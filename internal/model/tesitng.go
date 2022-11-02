package model

import (
	"testing"
	"time"
)

func TestSubscriber(t *testing.T) *Subscriber {
	return &Subscriber{
		Name:     "Bob",
		LastName: "Jhonson",
		BirthDay: time.Date(1960, 11, 02, 0, 0, 0, 0, time.UTC),
		Email:    "ex@test.org",
	}
}
