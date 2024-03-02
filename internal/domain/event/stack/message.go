package stack

import (
	"time"

	"github.com/taniko/blog/internal/domain/model/stack/vo"
)

type CreateMessageEvent struct {
	unimplemented
	id        vo.MessageID
	channelID vo.ChannelID
	content   vo.Content
	createdAt time.Time
}

func (e CreateMessageEvent) ID() vo.MessageID {
	return e.id
}

func (e CreateMessageEvent) ChannelID() vo.ChannelID {
	return e.channelID
}

func (e CreateMessageEvent) Content() vo.Content {
	return e.content
}

func (e CreateMessageEvent) CreatedAt() time.Time {
	return e.createdAt
}
