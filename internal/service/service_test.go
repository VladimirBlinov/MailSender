package service_test

import (
	"fmt"
	"testing"

	"github.com/VladimirBlinov/MailSender/internal/service"
	"github.com/VladimirBlinov/MailSender/internal/store/filestore"
	"github.com/VladimirBlinov/MailSender/pkg/email/smtpsender"
	smtpmock "github.com/mocktools/go-smtp-mock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_ServiceRunBroadcast(t *testing.T) {

	server := smtpmock.New(smtpmock.ConfigurationAttr{
		LogToStdout:       true,
		LogServerActivity: true,
	})

	if err := server.Start(); err != nil {
		fmt.Println(err)
	}

	hostAddress, portNumber := "127.0.0.1", server.PortNumber
	sender, err := smtpsender.NewSMTPSender("test@test.org", "", hostAddress, portNumber)
	if err != nil {
		fmt.Println(err)
	}

	if err := server.Start(); err != nil {
		fmt.Println(err)
	}

	testCases := []struct {
		name              string
		broadcastListPath string
		emailTempl        string
		subject           string
		delay             int
	}{
		{
			name:              "valid",
			broadcastListPath: "broadcastList_test.txt",
			emailTempl:        "template.html",
			subject:           "test",
			delay:             2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := logrus.New()
			store := filestore.NewStore(tc.broadcastListPath, logger)
			srvc := service.NewService(store, sender, logger)

			assert.NoError(t, srvc.RunBroadcast(tc.emailTempl, tc.subject, tc.delay))
			server.Messages()
		})
	}

	if err := server.Stop(); err != nil {
		fmt.Println(err)
	}
}
