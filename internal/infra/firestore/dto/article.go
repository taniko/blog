package dto

import (
	"github.com/taniko/blog/internal/domain/event/article"
)

type CreateArticle struct {
	EventHeader EventHeader `firestore:"header"`
	ArticleID   string      `firestore:"id"`
	Author      string      `firestore:"author"`
	Title       string      `firestore:"title"`
	Body        string      `firestore:"body"`
}

func NewCreateArticle(e article.CreateEvent) CreateArticle {
	return CreateArticle{
		EventHeader: NewEventHeader(e.GetHeader()),
		ArticleID:   e.GetArticleID().String(),
		Author:      e.GetAuthorID().String(),
		Title:       string(e.GetPayload().Title),
		Body:        string(e.GetPayload().Body),
	}
}
