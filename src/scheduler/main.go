package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaBrokers = []string{os.Getenv("BROKER_URL")}
	topics       = strings.Split(os.Getenv("TOPICS"), ",")
	//groupID       = ""
	w1 = &kafka.Writer{
		Addr:         kafka.TCP(kafkaBrokers[0]),
		Topic:        topics[0],
		RequiredAcks: kafka.RequireAll,
	}
	w2 = &kafka.Writer{
		Addr:         kafka.TCP(kafkaBrokers[0]),
		Topic:        topics[1],
		RequiredAcks: kafka.RequireAll,
	}
)

func task1() {
	// Add your task logic here
	currentTime := time.Now()
	message := fmt.Sprintf("Message: Scheduled Task 1: %s", currentTime)
	// Send Kafka message
	err := w1.WriteMessages(context.Background(), kafka.Message{Value: []byte(message)})
	if err != nil {
		log.Fatal("Error writing Kafka message:", err)
	}
}

func task2() {
	// Add your task logic here
	currentTime := time.Now()
	message := fmt.Sprintf("Message: Scheduled Task 2: %s", currentTime)

	// Send Kafka message
	err := w2.WriteMessages(context.Background(), kafka.Message{Value: []byte(message)})
	if err != nil {
		log.Fatal("Error writing Kafka message:", err)
	}
}

func main() {
	c := cron.New()

	_, err := c.AddFunc("*/1 * * * *", task1) // Run task1 every 5 minutes
	if err != nil {
		log.Fatal("Error adding task1 to cron:", err)
	}

	_, err = c.AddFunc("*/2 * * * *", task2) // Run task2 every day at midnight
	if err != nil {
		log.Fatal("Error adding task2 to cron:", err)
	}

	c.Start()

	select {}
}
