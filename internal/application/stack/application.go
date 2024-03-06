package stack

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/taniko/blog/internal/application/stack/command"
	"github.com/taniko/blog/internal/domain/event"
	"github.com/taniko/blog/internal/domain/model/post"
	"github.com/taniko/blog/internal/domain/model/stack"
	"github.com/taniko/blog/internal/domain/model/stack/vo"
	"github.com/taniko/blog/internal/domain/model/user"
	"github.com/taniko/blog/internal/domain/repository"
	"github.com/taniko/blog/internal/utils/ctxutil"
)

type Application interface {
	// CreateChannel チャンネルを作成する
	CreateChannel(ctx context.Context, cmd command.CreateChannel) error

	// CreateMessage メッセージを作成する
	CreateMessage(ctx context.Context, cmd command.CreateMessage) error
}

type application struct {
	stackRepository repository.Stack
}

func New(stackRepository repository.Stack) Application {
	return &application{
		stackRepository: stackRepository,
	}
}

func (a application) CreateChannel(ctx context.Context, cmd command.CreateChannel) error {
	userID, ok := ctxutil.GetUserID(ctx)
	if !ok {
		return user.ErrUnauthenticated
	}

	author := post.AuthorID(userID)

	_, e, err := stack.CreateChannel(author, vo.Name(cmd.Name), vo.Description(cmd.Description))
	if err != nil {
		return errors.Wrap(err, "create channel")
	}
	if err := a.stackRepository.Save(ctx, []event.Event{e}); err != nil {
		return errors.Wrap(err, "save event")
	}
	return nil
}

func (a application) CreateMessage(ctx context.Context, cmd command.CreateMessage) error {
	//TODO implement me
	panic("implement me")
}
