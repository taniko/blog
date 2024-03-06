package stack

import (
	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/model/post"
	"github.com/taniko/blog/internal/domain/model/stack/vo"
)

const NameCreateChannelEvent event.Name = "stack.v1.createChannel"

type CreateChannelEvent interface {
	event.Event

	// ChannelID チャンネルIDを取得
	ChannelID() vo.ChannelID

	// Name チャンネル名を取得
	Name() vo.Name

	// AuthorID 作成者IDを取得
	AuthorID() post.AuthorID

	// Description 説明を取得
	Description() vo.Description
}

type createChannelEvent struct {
	channelID   vo.ChannelID
	name        vo.Name
	authorID    post.AuthorID
	description vo.Description
	header      event.Header
}

func NewCreateChannelEvent(header event.Header, channelID vo.ChannelID, name vo.Name, authorID post.AuthorID, description vo.Description) CreateChannelEvent {
	return createChannelEvent{
		header:      header,
		channelID:   channelID,
		name:        name,
		authorID:    authorID,
		description: description,
	}
}

func (e createChannelEvent) GetEventHeader() event.Header {
	return e.header
}

func (e createChannelEvent) GetEventName() event.Name {
	return NameCreateChannelEvent
}

func (e createChannelEvent) ChannelID() vo.ChannelID {
	return e.channelID
}

func (e createChannelEvent) Name() vo.Name {
	return e.name
}

func (e createChannelEvent) AuthorID() post.AuthorID {
	return e.authorID
}

func (e createChannelEvent) Description() vo.Description {
	return e.description
}
