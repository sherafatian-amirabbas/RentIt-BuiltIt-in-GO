package kafka_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cs-ut-ee/project-group-3/pkg/domain"
	"github.com/segmentio/kafka-go"
)

func kafkaUrl() string {
	url, success := os.LookupEnv("kafkaUrl")
	if !success {
		url = "localhost:9092"
	}

	return url
}

func ApiUrl() string {
	url, success := os.LookupEnv("apiUrl")
	if !success {
		url = "http://localhost:8081"
	}

	return url
}

func customerId() string {
	customerId, success := os.LookupEnv("customerId")
	if !success {
		customerId = "customer-2"
	}

	return customerId
}

func TestReceivingValidInvoice(t *testing.T) {
	topic := customerId()
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaUrl(), topic, 0)
	if err != nil {
		log.Fatal(err)
	}

	invoice := domain.Invoice{
		ID:              11211,
		PurchaseOrderID: 1,
		Amount:          150.0,
		Regulator:       "",
		Comment:         "",
		Status:          3,
	}

	jsonMsg, _ := json.Marshal(invoice)

	msg := kafka.Message{
		Topic: topic,
		Key:   []byte(fmt.Sprintf("Invoice-%d", 11)),
		Value: []byte(jsonMsg),
	}

	_, err = conn.WriteMessages(msg)
	if err != nil {
		t.Error(err.Error())
	}

	// We don't have really good way how to test it because sending into kafka doesn't send respond
	// And because requirement is that when invoice is received info is written into stdout
	// There isn't really possibility to check api's stdout from test. Also it takes time for server
	// to handle new messages from kafka. Receiving invoice is goroutine and we don't know exact moment when invoice is receieved from kafka
	// Because of that we only check if sending to kafka is successful. This requirement should be ideally checked with unit test
	// Because then we know exact moment when invoice is received. But unit tests aren't required by project document
}
