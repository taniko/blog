package stack

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniko/blog/internal/domain/event"
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
		return Channel{}, nil, ErrEmptyName
	} else if description == "" {
		return Channel{}, nil, ErrEmptyDescription
	}
	id := vo.ChannelID(uuid.New().String())
	createdAt := time.Now()
	channel := Channel{
		id:          id,
		name:        name,
		description: description,
		createdAt:   createdAt,
	}
	e := stack.NewCreateChannelEvent(
		event.NewHeader(event.IssueID(), createdAt, 1),
		id,
		name,
		author,
		description,
	)
	return channel, e, nil
}
