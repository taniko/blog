package article

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/samber/mo"
	event "github.com/taniko/blog/internal/domain/event/article"
	"github.com/taniko/blog/internal/domain/model/article"
	"github.com/taniko/blog/internal/domain/model/article/vo"
	"github.com/taniko/blog/internal/domain/repository"
	"github.com/taniko/blog/internal/infra/firestore/dto"
)

const articleCollection = "articles"

type articleStore struct {
	c *firestore.Client
}

func NewArticleStore(c *firestore.Client) repository.Article {
	return &articleStore{
		c: c,
	}
}
func (a articleStore) Store(ctx context.Context, e event.Event) error {
	var data any
	switch e := e.(type) {
	case event.CreateEvent:
		data = dto.NewCreateArticle(e)
	default:
		return errors.New("unknown event")
	}
	_, err := a.lookupArticleDoc(e).
		Collection("events").
		Doc(e.GetEventHeader().GetVersion().String()).
		Set(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (a articleStore) Find(ctx context.Context, id vo.ID) mo.Result[article.Article] {
	panic("implement me")
}

// lookupArticleDoc articleのドキュメントを取得する
func (a articleStore) lookupArticleDoc(e event.Event) *firestore.DocumentRef {
	return a.c.Collection(articleCollection).
		Doc(e.GetAuthorID().String()).
		Collection("articles").
		Doc(e.GetArticleID().String())
}
