package filestore_test

import (
	"testing"

	"github.com/VladimirBlinov/MailSender/internal/store/filestore"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_SRGetSubscribers(t *testing.T) {
	logger := logrus.New()
	store := filestore.NewStore(broadcastListPath, logger)
	s, _ := store.Subscriber().GetSubscribers()
	assert.True(t, len(s) > 0)
}
