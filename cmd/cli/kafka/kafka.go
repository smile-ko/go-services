package kafka

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

// Consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers, // []string{"localhost:9092"}
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

// DebeziumMessage represents the structure of Debezium CDC events
type DebeziumMessage struct {
	Payload struct {
		Before map[string]interface{} `json:"before"`
		After  map[string]interface{} `json:"after"`
		Source struct {
			Table string `json:"table"`
			Op    string `json:"op"` // c: create, u: update, d: delete
		} `json:"source"`
	} `json:"payload"`
}

// ConsumeDebeziumMessages starts consuming CDC events from Kafka
func ConsumeDebeziumMessages(kafkaURL, topic, groupID string) {
	reader := getKafkaReader(kafkaURL, topic, groupID)
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var debeziumMsg DebeziumMessage
		if err := json.Unmarshal(msg.Value, &debeziumMsg); err != nil {
			log.Printf("Error unmarshaling message: %v", err)
			continue
		}

		// Log the change
		logDatabaseChange(debeziumMsg)
	}
}

// logDatabaseChange logs database changes from Debezium events
func logDatabaseChange(msg DebeziumMessage) {
	operation := msg.Payload.Source.Op
	table := msg.Payload.Source.Table

	switch operation {
	case "c":
		log.Printf("INSERT on table %s: %v", table, msg.Payload.After)
	case "u":
		log.Printf("UPDATE on table %s:\nBefore: %v\nAfter: %v",
			table, msg.Payload.Before, msg.Payload.After)
	case "d":
		log.Printf("DELETE on table %s: %v", table, msg.Payload.Before)
	}
}
