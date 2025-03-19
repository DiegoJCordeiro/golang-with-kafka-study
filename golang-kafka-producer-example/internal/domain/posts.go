package domain

import "time"

type Posts struct {
	uuid      string
	message   string
	readAt    time.Time
	deletedAt time.Time
}

func NewPosts(uuid string, message string, readAt time.Time, deletedAt time.Time) *Posts {
	return &Posts{uuid, message, readAt, deletedAt}
}

func (p *Posts) UUID() string {
	return p.uuid
}

func (p *Posts) Message() string {
	return p.message
}

func (p *Posts) ReadAt() time.Time {
	return p.readAt
}

func (p *Posts) DeletedAt() time.Time {
	return p.deletedAt
}
