//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE
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
