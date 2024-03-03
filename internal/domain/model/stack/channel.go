package stack

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniko/blog/internal/domain/event/stack"
	"github.com/taniko/blog/internal/domain/model/post"
	"github.com/taniko/blog/internal/domain/model/stack/vo"
)

type Channel struct {
	id          vo.ChannelID
	name        vo.Name
	description vo.Description
	createdAt   time.Time
}

func CreateChannel(author post.AuthorID, name vo.Name, description vo.Description) (Channel, stack.CreateChannelEvent, error) {
	if name == "" {
		return Channel{}, stack.CreateChannelEvent{}, ErrEmptyName
	} else if description == "" {
		return Channel{}, stack.CreateChannelEvent{}, ErrEmptyDescription
	}
	id := vo.ChannelID(uuid.New().String())
	createdAt := time.Now()
	channel := Channel{
		id:          id,
		name:        name,
		description: description,
		createdAt:   createdAt,
	}
	event := stack.NewCreateChannelEvent(id, name, author, description, createdAt)
	return channel, event, nil
}
