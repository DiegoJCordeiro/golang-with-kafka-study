package usecase

import (
	"errors"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/kafka_client"
)

type CreatePostsUseCase interface {
	Execute(posts []string) error
	isValid() (bool, error)
}

type CreatePostsUseCaseImpl struct {
	kafkaClient kafka_client.KafkaClient
}

func NewCreatePostsUseCase(kafkaClient kafka_client.KafkaClient) CreatePostsUseCase {
	return &CreatePostsUseCaseImpl{
		kafkaClient: kafkaClient,
	}
}

func (createPostsUseCase *CreatePostsUseCaseImpl) Execute(posts []string) error {

	if ok, err := createPostsUseCase.isValid(); !ok {
		return err
	}

	for _, post := range posts {
		createPostsUseCase.kafkaClient.AddMessage("tp_notification_posts", []byte(post))
	}

	if err := createPostsUseCase.kafkaClient.Produce(); err != nil {
		return err
	}

	return nil
}

func (createPostsUseCase *CreatePostsUseCaseImpl) isValid() (bool, error) {

	if createPostsUseCase.kafkaClient == nil {
		return false, errors.New("kafkaClient is nil")
	}

	return true, nil
}
