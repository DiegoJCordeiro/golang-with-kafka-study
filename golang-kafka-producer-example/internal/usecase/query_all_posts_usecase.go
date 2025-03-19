package usecase

import (
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/dto"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/database/repository"
)

type QueryAllPostsUseCase interface {
	Execute(limit int64, offset int64) (dto.QueryPostsDTO, error)
	isValid() (bool, error)
}

type QueryAllPostsUseCaseImpl struct {
	postRepository repository.PostRepository
}

func NewQueryAllPostsUseCase(postRepository repository.PostRepository) QueryAllPostsUseCase {
	return &QueryAllPostsUseCaseImpl{
		postRepository: postRepository,
	}
}

func (queryUseCase *QueryAllPostsUseCaseImpl) Execute(limit int64, offset int64) (dto.QueryPostsDTO, error) {

	queryPostsDTO, err := queryUseCase.postRepository.QueryAll(limit, offset)

	if err != nil {
		return dto.QueryPostsDTO{}, err
	}

	return queryPostsDTO, nil
}

func (queryUseCase *QueryAllPostsUseCaseImpl) isValid() (bool, error) {
	return true, nil
}
