package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	brokerAddress = os.Getenv("BROKER_URL")
	topics        = strings.Split(os.Getenv("TOPICS"), ",")
	//groupID       = ""
)

func main() {
	config := kafka.ReaderConfig{
		Brokers:     []string{brokerAddress},
		GroupID:     "consumer-A",
		GroupTopics: topics,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	ctx := context.Background()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Starting Kafka consumer...")
	for {
		select {
		case <-ctx.Done():
			return
		case sig := <-signals:
			fmt.Printf("Received signal: %s\n", sig)
			return
		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Println("Error reading Kafka message:", err)
				continue
			}
			fmt.Printf("Received message: %s\n", string(msg.Value))
		}
	}
}
