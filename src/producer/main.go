package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaBrokers = []string{os.Getenv("BROKER_URL")}
	topics       = strings.Split(os.Getenv("TOPICS"), ",")
)

func sendKafkaMessage(message string, topic string) {
	// Kafka configuration
	w := &kafka.Writer{
		Addr:         kafka.TCP(kafkaBrokers[0]),
		Topic:        topic,
		RequiredAcks: kafka.RequireAll,
	}
	w.AllowAutoTopicCreation = true
	w.Topic = topic
	err := w.WriteMessages(context.Background(), kafka.Message{Value: []byte(message)})
	if err != nil {
		log.Fatal("Error writing Kafka message:", err)
	}
	defer w.Close()
}

func main() {
	for {
		for _, topic := range topics {
			// Add your task logic here
			fmt.Println("topic: ", topic)
			currentTime := time.Now()
			message := fmt.Sprintf("Message: Producer sent to topic: %s at: %s", topic, currentTime)
			// Send Kafka message
			sendKafkaMessage(message, topic)
		}
	}
}
