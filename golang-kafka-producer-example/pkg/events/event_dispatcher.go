package events

type EventDispatcher interface {
	Register(topicName string, handler EventHandler) error
	RemoveTopic(topicName string) error
	RemoveHandler(topicName string, handler EventHandler) error
	HasTopic(topicName string) bool
	HasHandler(topicName string, handler EventHandler) bool
	Dispatch(topicName string, event Event) error
	Clear() error
}
