package article

import (
	"github.com/google/uuid"
	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/model/article/vo"
)

const NameCreate event.Name = "article.v1.create"

type (
	CreateEvent interface {
		Event
		// GetPayload ペイロードを取得
		GetPayload() CrateEventPayload
	}
	CrateEventPayload struct {
		ID     vo.ID
		Author vo.AuthorID
		Title  vo.Title
		Body   vo.Body
	}
	createEvent struct {
		header  event.Header
		payload CrateEventPayload
	}
)

var _ Event = (*createEvent)(nil)

func (c createEvent) GetEventHeader() event.Header {
	return c.header
}

func (c createEvent) GetEventName() event.Name {
	return NameCreate
}

func (c createEvent) GetAuthorID() vo.AuthorID {
	return c.GetPayload().Author
}

func (c createEvent) GetArticleID() vo.ID {
	return c.GetPayload().ID
}

func NewCreateEvent(header event.Header, author vo.AuthorID, title vo.Title, body vo.Body) CreateEvent {
	return createEvent{
		header: header,
		payload: CrateEventPayload{
			ID:     vo.ID(uuid.NewString()),
			Author: author,
			Title:  title,
			Body:   body,
		},
	}
}

func (c createEvent) GetPayload() CrateEventPayload {
	return c.payload
}
