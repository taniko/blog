package article

import (
	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/model/article/vo"
)

type Event interface {
	event.Event

	// GetAuthorID 作成者IDを取得
	GetAuthorID() vo.AuthorID

	// GetArticleID 記事IDを取得
	GetArticleID() vo.ID
}
