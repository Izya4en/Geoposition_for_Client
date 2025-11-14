package kafka

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	asyncProducer sarama.AsyncProducer
}

func NewProducer(broker string) *Producer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Kafka producer error: %v", err)
	}

	go func() {
		for err := range producer.Errors() {
			log.Printf("Kafka error: %v", err)
		}
	}()

	return &Producer{asyncProducer: producer}
}

func (p *Producer) SendMessage(topic, message string) {
	msg := &sarama.ProducerMessage{Topic: topic, Value: sarama.StringEncoder(message)}
	p.asyncProducer.Input() <- msg
}
