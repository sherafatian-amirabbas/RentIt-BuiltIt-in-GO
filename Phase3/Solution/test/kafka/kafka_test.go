package kafkatest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	kafka "github.com/segmentio/kafka-go"
)

func kafkaUrl() string {
	url, success := os.LookupEnv("kafkaUrl")
	if !success {
		url = "localhost:9092"
	}

	return url
}

func apiUrl() string {
	url, success := os.LookupEnv("httpUrl")
	if !success {
		url = "http://localhost:8081"
	}

	return url
}

func TestKafka(t *testing.T) {
	_, err := http.Post(apiUrl()+"/customers/2/invoices/sendReminder", "application/json", nil)
	if err != nil {
		t.Error("TestKafka: Problem sending reminder.", err)
		return
	}

	reader := getKafkaReader(kafkaUrl(), "customerds-2", "1")

	defer reader.Close()

	fmt.Println("start consuming ... !!")

	m, err := reader.ReadMessage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	if string(m.Key) != "Invoice-1" {
		t.Error("Expected: " + "Invoice-1" + "\nActual: " + string(m.Key))
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
