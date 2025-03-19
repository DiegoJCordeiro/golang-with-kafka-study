package repository

import (
	"context"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/dto"
	"github.com/DiegoJCordeiro/golang-with-kafka-study/golang-kafka-producer/internal/infra/database/sqlc"
)

type PostRepository interface {
	QueryAll(limit int64, offset int64) (dto.QueryPostsDTO, error)
}

type PostRepositoryImpl struct {
	db      *sqlc.DBTX
	queries *sqlc.Queries
}

func NewPostRepository(db sqlc.DBTX) *PostRepositoryImpl {
	return &PostRepositoryImpl{
		db:      &db,
		queries: sqlc.New(db),
	}
}

func (postsRepository *PostRepositoryImpl) QueryAll(limit int64, offset int64) (dto.QueryPostsDTO, error) {

	ctx := context.Background()

	posts, err := postsRepository.queries.QueryAllPosts(ctx, sqlc.QueryAllPostsParams{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return dto.QueryPostsDTO{}, err
	}

	postsQueried := make([]dto.Posts, len(posts))

	for _, post := range posts {
		postsQueried = append(postsQueried, dto.Posts{
			Message:   post.Message.String,
			ReadAt:    post.ReadAt.Time,
			DeletedAt: post.DeletedAt.Time,
		})
	}

	return dto.QueryPostsDTO{
		Messages: postsQueried,
	}, nil
}
