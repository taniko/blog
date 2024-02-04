package article

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/model/article/vo"
)

var _ Event = (*CreateEvent)(nil)

type (
	CreateEvent struct {
		unimplemented
		header  event.Header
		payload CrateEventPayload
	}

	CrateEventPayload struct {
		ID     vo.ID
		Author vo.AuthorID
		Title  vo.Title
		Body   vo.Body
	}
)

func (c CreateEvent) GetArticleID() vo.ID {
	return c.GetPayload().ID
}

func NewCreateEvent(author vo.AuthorID, title vo.Title, body vo.Body) CreateEvent {
	return CreateEvent{
		header: event.NewHeader(
			event.IssueID(),
			time.Now(),
			event.Version(1),
			CreateName,
		),
		payload: CrateEventPayload{
			ID:     vo.ID(uuid.NewString()),
			Author: author,
			Title:  title,
			Body:   body,
		},
	}
}

func (c CreateEvent) GetHeader() event.Header {
	return c.header
}

func (c CreateEvent) GetPayload() CrateEventPayload {
	return c.payload
}
