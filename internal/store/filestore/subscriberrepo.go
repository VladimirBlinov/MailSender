package filestore

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/VladimirBlinov/MailSender/internal/model"
)

type SubscriberRepo struct {
	store *Store
}

func (sr *SubscriberRepo) GetSubscribers() ([]*model.Subscriber, error) {
	file, err := os.Open(sr.store.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var subscribers []*model.Subscriber

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return nil, err
		}

		sl := strings.Split(string(line), ";")

		s := &model.Subscriber{
			Name:     sl[0],
			LastName: sl[1],
			Email:    sl[3],
		}

		bdParsed, err := time.Parse("2006-01-02", sl[2])
		s.BirthDay = bdParsed

		if err := s.Validate(); err != nil {
			log.Printf("read file line error: %v", err)
			continue
		}

		subscribers = append(subscribers, s)
	}
	return subscribers, nil
}
