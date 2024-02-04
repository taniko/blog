package article

import (
	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/model/article/vo"
)

type Event interface {
	internal()
	event.Event
	GetArticleID() vo.ID
}

type unimplemented struct{}

func (unimplemented) internal() {}
