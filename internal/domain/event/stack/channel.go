package stack

import (
	"time"

	"github.com/taniko/blog/internal/domain/model/post"
	"github.com/taniko/blog/internal/domain/model/stack/vo"
)

type CreateChannelEvent struct {
	unimplemented
	id          vo.ChannelID
	name        vo.Name
	authorID    post.AuthorID
	description vo.Description
	createdAt   time.Time
}

func NewCreateChannelEvent(id vo.ChannelID, name vo.Name, authorID post.AuthorID, description vo.Description, createdAt time.Time) CreateChannelEvent {
	return CreateChannelEvent{
		id:          id,
		name:        name,
		authorID:    authorID,
		description: description,
		createdAt:   createdAt,
	}
}

func (e CreateChannelEvent) ID() vo.ChannelID {
	return e.id
}

func (e CreateChannelEvent) Name() vo.Name {
	return e.name
}

func (e CreateChannelEvent) AuthorID() post.AuthorID {
	return e.authorID
}

func (e CreateChannelEvent) Description() vo.Description {
	return e.description
}

func (e CreateChannelEvent) CreatedAt() time.Time {
	return e.createdAt
}
