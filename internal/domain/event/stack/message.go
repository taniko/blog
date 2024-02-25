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
