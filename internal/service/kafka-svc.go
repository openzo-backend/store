package service

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	// "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaService interface {
	Consume(topic string, handler func(*kafka.Message) error) error
	Produce(topic string, message []byte) error
}

type kafkaServiceImpl struct {
	consumer   *kafka.Consumer
	producer   *kafka.Producer
	retryCount int
	retryDelay time.Duration
	topics     []string
}

func NewKafkaService(consumerConf, producerConf kafka.ConfigMap, retryCount int, retryDelay time.Duration, topics []string) (KafkaService, error) {
	consumer, err := kafka.NewConsumer(&consumerConf)
	if err != nil {
		return nil, fmt.Errorf("error creating consumer: %v", err)
	}
	fmt.Println("Connected to Kafka cluster.")

	producer, err := kafka.NewProducer(&producerConf)
	if err != nil {
		return nil, fmt.Errorf("error creating producer: %v", err)
	}

	return &kafkaServiceImpl{
		consumer:   consumer,
		producer:   producer,
		retryCount: retryCount,
		retryDelay: retryDelay,
	}, nil
}

func (ks *kafkaServiceImpl) Consume(topic string, handler func(*kafka.Message) error) error {
	err := ks.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return fmt.Errorf("error subscribing to topic: %v", err)
	}

	run := true
	for run {
		msg, err := ks.consumer.ReadMessage(-1)
		if err == nil {
			if err := handler(msg); err != nil {
				fmt.Printf("Error processing message: %v\n", err)
			}
		} else {
			fmt.Printf("Error consuming message: %v\n", err)
			if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrAllBrokersDown {
				fmt.Println("All brokers are down. Attempting to reconnect...")
				return ks.retryConnection()
			}
		}
	}

	return nil
}

func (ks *kafkaServiceImpl) Produce(topic string, message []byte) error {
	return ks.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)
}

func (ks *kafkaServiceImpl) retryConnection() error {
	for i := 0; i < ks.retryCount; i++ {
		time.Sleep(ks.retryDelay)

		err := ks.consumer.SubscribeTopics(ks.topics, nil)
		if err == nil {
			fmt.Println("Reconnected to Kafka.")
			return nil
		}
		fmt.Fprintf(os.Stderr, "Reconnection attempt %d failed: %v\n", i+1, err)
		ks.retryDelay *= 2
	}

	return fmt.Errorf("failed to reconnect after %d attempts", ks.retryCount)
}
