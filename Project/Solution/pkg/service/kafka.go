package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/cs-ut-ee/project-group-3/pkg/domain"
	kafka "github.com/segmentio/kafka-go"
)

func kafkaUrl() string {
	url, success := os.LookupEnv("kafkaUrl")
	if !success {
		url = "localhost:9092"
	}

	return url
}

func customerId() (string, bool) {
	url, success := os.LookupEnv("customerId")
	if !success {
		return "", false
	}

	return url, true
}

func (service Service) ReceiveInvoiceJob() {
	customerId, success := customerId()
	if !success {
		return
	}

	reader := getKafkaReader(kafkaUrl(), customerId, "1")
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		invoice := &domain.Invoice{}

		err = json.Unmarshal(msg.Value, invoice)
		if err != nil {
			log.Println(err)
		}

		invoice, err = service.Repository.NewInvoice(invoice)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("%v", invoice)
		}
	}
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
