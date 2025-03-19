package events

import (
	"errors"
)

type EventKafka struct {
	uuid      *string
	name      *string
	message   *[][]byte
	timestamp *string
}

func NewEventKafkaProducer(name, timestamp, uuid string, message [][]byte) Event {
	return &EventKafka{
		uuid:      &uuid,
		name:      &name,
		message:   &message,
		timestamp: &timestamp,
	}
}

func (e *EventKafka) GetUUID() *string {
	return e.uuid
}

func (e *EventKafka) GetName() *string {
	return e.name
}

func (e *EventKafka) GetMessages() *[][]byte {
	return e.message
}

func (e *EventKafka) GetTimestamp() *string {
	return e.timestamp
}

func (e *EventKafka) IsValid() error {

	if *e.name == "" {
		return errors.New("topicName can't be empty")
	} else if *e.timestamp == "" {
		return errors.New("timestamp can't be empty")
	} else if *e.message == nil {
		return errors.New("message can't be empty")
	} else {
		return nil
	}
}

func (e *EventKafka) AddMessage(newMessage []byte) {
	sizeSlice := len(*e.GetMessages())
	var newSlice [][]byte
	if sizeSlice > 0 {
		newSlice = make([][]byte, 0, sizeSlice+1)
		for index, message := range *e.GetMessages() {
			if index == 0 {
				newSlice = append(newSlice, newMessage)
			}
			newSlice = append(newSlice, message)
		}
	} else {
		newSlice = make([][]byte, 0, 1)
		newSlice = append(newSlice, newMessage)
	}
	e.message = &newSlice
}
