package testsender

import (
	"github.com/VladimirBlinov/MailSender/pkg/email"
	"github.com/stretchr/testify/mock"
)

type EmailSender struct {
	mock.Mock
}

func (m *EmailSender) Send(inp email.SendEmailInput) error {
	args := m.Called(inp)

	return args.Error(0)
}
