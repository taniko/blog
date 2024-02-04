//go:generate mockgen -source=article.go -destination=mock/article.go -package=article_mock
package repository

import (
	"context"

	"github.com/samber/mo"
	event "github.com/taniko/blog/internal/domain/event/article"
	"github.com/taniko/blog/internal/domain/model/article"
	"github.com/taniko/blog/internal/domain/model/article/vo"
)

type Article interface {
	Store(ctx context.Context, e event.Event) error
	Find(ctx context.Context, id vo.ID) mo.Result[article.Article]
}
