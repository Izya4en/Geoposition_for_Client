package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

func StartConsumer(broker, topic string) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{broker}, config)
	if err != nil {
		panic(fmt.Sprintf("Error creating consumer: %v", err))
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(fmt.Sprintf("Error starting partition consumer: %v", err))
	}
	defer partitionConsumer.Close()

	fmt.Println("Kafka consumer listening on topic:", topic)
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Received message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("Error: %v\n", err)
		case <-context.Background().Done():
			return
		}
	}
}
