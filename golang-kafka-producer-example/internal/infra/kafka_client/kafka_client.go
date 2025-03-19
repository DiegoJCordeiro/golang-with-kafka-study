package kafka_client

import (
	"errors"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/pkg/events"
	"github.com/google/uuid"
	"log"
	"time"
)

type KafkaClient interface {
	AddMessage(topic string, message []byte)
	DeleteMessage(topic string) error
	Produce() error
}

type KafkaClientImpl struct {
	bootstrapServer string
	eventDispatcher events.EventDispatcherKafka
	events          map[string]events.Event
}

func NewKafkaClient(bootstrapServer string) *KafkaClientImpl {
	return &KafkaClientImpl{
		bootstrapServer: bootstrapServer,
		eventDispatcher: *events.NewEventDispatcherKafka(),
		events:          make(map[string]events.Event),
	}
}

func (kafkaClient *KafkaClientImpl) AddMessage(topic string, message []byte) {

	uuidEvent, _ := uuid.NewUUID()

	if _, ok := kafkaClient.events[topic]; !ok {
		newSlice := make([][]byte, 0)
		kafkaClient.events[topic] = events.NewEventKafkaProducer(topic, uuidEvent.String(), time.Now().String(), newSlice)
	}

	kafkaClient.events[topic].AddMessage(message)
}

func (kafkaClient *KafkaClientImpl) DeleteMessage(topic string) error {

	if _, ok := kafkaClient.events[topic]; !ok {
		return errors.New("this topic not exist")
	}

	delete(kafkaClient.events, topic)

	return nil
}

func (kafkaClient *KafkaClientImpl) Produce() error {

	eventHandler, err := events.NewEventHandlerKafka(kafkaClient.bootstrapServer)

	if err != nil {
		errors.New("unable create the event handler")
	}

	if len(kafkaClient.events) > 0 {
		for index, event := range kafkaClient.events {
			if err := kafkaClient.eventDispatcher.Register(index, eventHandler); err != nil {
				log.Fatal(err.Error())
			}
			err = kafkaClient.eventDispatcher.Dispatch(index, event)

			if err != nil {
				log.Fatal(err.Error())
				return err
			}
			log.Println("message sent to the topic")
		}

		return nil
	}

	return errors.New("not has events to produce")
}
