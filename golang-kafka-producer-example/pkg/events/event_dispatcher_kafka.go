package events

import (
	"errors"
)

type EventDispatcherKafka struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcherKafka() *EventDispatcherKafka {
	return &EventDispatcherKafka{
		handlers: make(map[string][]EventHandler),
	}
}

func (eventDispatcher *EventDispatcherKafka) Register(topicName string, handler EventHandler) error {

	if _, ok := eventDispatcher.handlers[topicName]; ok {
		for _, handlerIterated := range eventDispatcher.handlers[topicName] {
			if handlerIterated == handler {
				return errors.New("event dispatcher already has the handler registered to this topic" + topicName)
			}
		}
	}

	eventDispatcher.handlers[topicName] = append(eventDispatcher.handlers[topicName], handler)

	return nil
}

func (eventDispatcher *EventDispatcherKafka) RemoveTopic(topicName string) error {

	if _, ok := eventDispatcher.handlers[topicName]; !ok {
		return errors.New("event dispatcher has not the topic name registered" + topicName)
	}

	delete(eventDispatcher.handlers, topicName)

	return nil
}

func (eventDispatcher *EventDispatcherKafka) RemoveHandler(topicName string, handler EventHandler) error {

	if handlers, ok := eventDispatcher.handlers[topicName]; ok {
		for index, handlerIterated := range handlers {
			if handlerIterated == handler {
				eventDispatcher.handlers[topicName] = append(handlers[:index], handlers[index+1:]...)
			}
		}
	}

	return nil
}

func (eventDispatcher *EventDispatcherKafka) HasTopic(topicName string) bool {

	if _, ok := eventDispatcher.handlers[topicName]; ok {
		return true
	}

	return false
}

func (eventDispatcher *EventDispatcherKafka) HasHandler(topicName string, handler EventHandler) bool {

	if handlers, ok := eventDispatcher.handlers[topicName]; ok {
		for _, handlerIterated := range handlers {
			if handlerIterated == handler {
				return true
			}
		}
	}

	return false
}

func (eventDispatcher *EventDispatcherKafka) Dispatch(topicName string, event Event) error {

	if handlers, ok := eventDispatcher.handlers[topicName]; ok {
		for _, handlerIterated := range handlers {
			if err := handlerIterated.Handle(event); err != nil {
				return err
			}
		}
	}

	return nil
}

func (eventDispatcher *EventDispatcherKafka) Clear() error {

	for key := range eventDispatcher.handlers {
		delete(eventDispatcher.handlers, key)
	}

	return nil
}
