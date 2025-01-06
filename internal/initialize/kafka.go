package initialize

import (
	"go-services/cmd/cli/kafka"
	"log"
)

func InitKafka() {
	// Kafka configuration
	kafkaURL := "kafka:9092"
	topic := "dbserver1.public.go_db_user"
	groupID := "cdc-group"

	// Start consuming messages in a goroutine
	go func() {
		log.Println("Starting Kafka CDC consumer...")
		kafka.ConsumeDebeziumMessages(kafkaURL, topic, groupID)
	}()
}
