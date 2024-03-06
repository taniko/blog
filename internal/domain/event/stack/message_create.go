//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE
package stack

import (
	"time"

	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/model/stack/vo"
)

const (
	NameCreateMessageEvent event.Name = "stack.v1.createMessageEvent"
)

type CreateMessageEvent interface {
	event.Event

	// ChannelID チャンネルIDを取得
	ChannelID() vo.ChannelID

	// MessageID メッセージIDを取得
	MessageID() vo.MessageID

	// Content コンテンツを取得
	Content() vo.Content

	// CreatedAt 作成日時を取得
	CreatedAt() time.Time
}

type createMessageEvent struct {
	header    event.Header
	messageID vo.MessageID
	channelID vo.ChannelID
	content   vo.Content
	createdAt time.Time
}

func NewCreateMessage(header event.Header, channelID vo.ChannelID, content vo.Content, createdAt time.Time) CreateMessageEvent {
	return createMessageEvent{
		header:    header,
		channelID: channelID,
		content:   content,
		createdAt: createdAt,
	}
}

func (c createMessageEvent) GetEventName() event.Name {
	return NameCreateMessageEvent
}

func (c createMessageEvent) ChannelID() vo.ChannelID {
	return c.channelID
}

func (c createMessageEvent) Content() vo.Content {
	return c.content
}

func (c createMessageEvent) CreatedAt() time.Time {
	return c.createdAt
}

func (c createMessageEvent) GetEventHeader() event.Header {
	return c.header
}

func (c createMessageEvent) MessageID() vo.MessageID {
	return c.messageID
}
