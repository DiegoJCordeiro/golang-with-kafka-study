package events

type Event interface {
	GetUUID() *string
	GetName() *string
	GetMessages() *[][]byte
	GetTimestamp() *string
	IsValid() error
	AddMessage(message []byte)
}
