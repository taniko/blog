package stack

import (
	"time"

	"github.com/taniko/blog/internal/domain/model/stack/vo"
)

type Message struct {
	ID        vo.MessageID
	Content   vo.Content
	CreatedAt time.Time
	ChannelID vo.ChannelID
}
