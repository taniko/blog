package article

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/taniko/blog/internal/application/article/command"
	"github.com/taniko/blog/internal/domain/model/account"
	"github.com/taniko/blog/internal/domain/model/article"
	"github.com/taniko/blog/internal/domain/model/article/vo"
	"github.com/taniko/blog/internal/domain/repository"
	"github.com/taniko/blog/internal/utils/ctxutil"
)

type Service interface {
	Create(ctx context.Context, cmd command.Create) error
}

type service struct {
	repo repository.Article
}

func New(repo repository.Article) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, cmd command.Create) error {
	userID, ok := ctxutil.Value[account.ID](ctx)
	if !ok {
		return errors.New("cannot get account id")
	}
	e := article.Create(vo.AuthorID(userID), vo.Title(cmd.Title), vo.Body(cmd.Body))
	if err := s.repo.Store(ctx, e); err != nil {
		return errors.Wrap(err, "store create article event")
	}
	return nil
}
