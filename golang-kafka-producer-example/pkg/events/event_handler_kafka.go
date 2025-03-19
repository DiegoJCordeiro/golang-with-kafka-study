package events

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"sync"
)

type EventHandlerKafkaImpl struct {
	producer  *kafka.Producer
	waitGroup *sync.WaitGroup
}

func NewEventHandlerKafka(bootstrapServer string) (EventHandler, error) {

	// "bootstrap.servers": "localhost:9092"
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	}

	producer, err := kafka.NewProducer(configMap)

	if err != nil {
		return nil, err
	}

	return &EventHandlerKafkaImpl{
		waitGroup: &sync.WaitGroup{},
		producer:  producer,
	}, nil
}

func (eventHandler *EventHandlerKafkaImpl) Handle(event Event) error {

	if messagesLength := len(*event.GetMessages()); messagesLength > 0 {
		eventHandler.waitGroup.Add(messagesLength)
	}

	for _, message := range *event.GetMessages() {
		err := eventHandler.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     event.GetName(),
				Partition: kafka.PartitionAny,
			},
			Value: message,
		}, nil)

		if err != nil {
			return err
		}
		eventHandler.waitGroup.Done()
	}

	eventHandler.producer.Flush(1000)

	return nil
}
