package model_test

import (
	"testing"
	"time"

	"github.com/VladimirBlinov/MailSender/internal/model"
	"github.com/stretchr/testify/assert"
)

func Test_SubscriberValidate(t *testing.T) {
	testCases := []struct {
		name    string
		s       func() *model.Subscriber
		isValid bool
	}{
		{
			name: "valid",
			s: func() *model.Subscriber {
				return model.TestSubscriber(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			s: func() *model.Subscriber {
				s := model.TestSubscriber(t)
				s.Email = ""
				return s
			},
			isValid: false,
		},
		{
			name: "short name",
			s: func() *model.Subscriber {
				s := model.TestSubscriber(t)
				s.Name = ""
				return s
			},
			isValid: false,
		},
		{
			name: "short lastname",
			s: func() *model.Subscriber {
				s := model.TestSubscriber(t)
				s.LastName = ""
				return s
			},
			isValid: false,
		},
		{
			name: "empty bd",
			s: func() *model.Subscriber {
				s := model.TestSubscriber(t)
				s.BirthDay = time.Time{}
				return s
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().Validate())
			} else {
				assert.Error(t, tc.s().Validate())
			}
		})
	}
}
