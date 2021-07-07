package service

import (
	"context"
	"fmt"
	"log"
	"os"

	kafka "github.com/segmentio/kafka-go"
)

func kafkaUrl() string {
	url, success := os.LookupEnv("kafkaUrl")
	if !success {
		url = "localhost:9092"
	}

	return url
}

func (service Service) SendReminderFor(customerID int64) error {
	topic := fmt.Sprintf("customerds-%d", customerID)
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaUrl(), topic, 0)
	if err != nil {
		log.Fatal(err)
	}

	invoices, err := service.Repository.GetUnPaidInvoicesFor(customerID)
	if err != nil {
		log.Fatal(err)
	}
	var msgs []kafka.Message
	for _, invoice := range invoices {
		msgs = append(msgs, kafka.Message{
			Topic: topic,
			Key:   []byte(fmt.Sprintf("Invoice-%d", invoice.ID)),
			Value: []byte(fmt.Sprintf("You have unpaid invoice.\nInvoice number: %d\nPrice: %f eur", invoice.ID, invoice.Price)),
		})
	}

	_, err = conn.WriteMessages(msgs...)
	if err != nil {
		println(err.Error())
	}

	return nil
}

func (service Service) SendReminders() error {
	customers, err := service.Repository.GetAllCustomers()
	if err != nil {
		log.Fatal(err)
	}
	for _, customerd := range customers {
		err = service.SendReminderFor(customerd.ID)
		if err != nil {
			println(err.Error())
		}
	}

	return nil
}
