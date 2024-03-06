package article

import (
	"time"

	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/event/article"
	"github.com/taniko/blog/internal/domain/model/article/vo"
)

type Article struct {
	id        vo.ID
	author    vo.AuthorID
	title     vo.Title
	body      vo.Body
	createdAt time.Time
	updatedAt time.Time
}

func New(id vo.ID, author vo.AuthorID, title vo.Title, body vo.Body) *Article {
	return &Article{
		id:        id,
		author:    author,
		title:     title,
		body:      body,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func Create(author vo.AuthorID, title vo.Title, body vo.Body) article.CreateEvent {
	return article.NewCreateEvent(event.NewHeader(event.IssueID(), time.Now(), 1), author, title, body)
}
