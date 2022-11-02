package filestore_test

import (
	"testing"

	"github.com/VladimirBlinov/MailSender/internal/store/filestore"
	"github.com/stretchr/testify/assert"
)

func Test_SRGetSubscribers(t *testing.T) {
	store := filestore.NewStore(broadcastListPath)
	s, _ := store.Subscriber().GetSubscribers()
	assert.True(t, len(s) > 0)
}
