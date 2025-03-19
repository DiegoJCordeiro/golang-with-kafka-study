package dto

import "time"

type Posts struct {
	Message   string    `json:"message"`
	ReadAt    time.Time `json:"readAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type CreatePostsDTO struct {
	Messages []string `json:"posts"`
}

type QueryPostsDTO struct {
	Messages []Posts `json:"posts"`
}
