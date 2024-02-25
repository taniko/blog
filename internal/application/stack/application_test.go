package stack

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/blog/internal/application/stack/command"
	"github.com/taniko/blog/internal/domain/event/stack"
	"github.com/taniko/blog/internal/domain/model/post"
	stack2 "github.com/taniko/blog/internal/domain/model/stack"
	vo2 "github.com/taniko/blog/internal/domain/model/stack/vo"
	"github.com/taniko/blog/internal/domain/model/user"
	"github.com/taniko/blog/internal/domain/model/user/vo"
	mock "github.com/taniko/blog/internal/domain/repository/mock"
	"github.com/taniko/blog/internal/utils/ctxutil"
	"go.uber.org/mock/gomock"
)

func TestApplication_CreateChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type test struct {
		author  vo.UserID
		name    string
		cmd     command.CreateChannel
		hasErr  bool
		wantErr error
	}

	tests := []test{
		{
			name:   "success",
			author: vo.UserID(uuid.NewString()),
			cmd: command.CreateChannel{
				Name:        "test",
				Description: uuid.NewString(),
			},
			hasErr: false,
		}, {
			name:   "unauthenticated",
			author: "",
			cmd: command.CreateChannel{
				Name:        "test",
				Description: uuid.NewString(),
			},
			wantErr: user.ErrUnauthenticated,
		}, {
			name:   "empty name",
			author: vo.UserID(uuid.NewString()),
			cmd: command.CreateChannel{
				Name:        "",
				Description: uuid.NewString(),
			},
			wantErr: stack2.ErrEmptyName,
		}, {
			name:   "empty description",
			author: vo.UserID(uuid.NewString()),
			cmd: command.CreateChannel{
				Name:        "test",
				Description: "",
			},
			wantErr: stack2.ErrEmptyDescription,
		},
	}

	var events []stack.Event
	repo := mock.NewMockStack(ctrl)
	repo.EXPECT().Save(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, es []stack.Event) error {
		events = es
		return nil
	}).AnyTimes()
	app := New(repo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			ctx := context.Background()
			if tt.author != "" {
				fmt.Println(ctxutil.GetUserID(ctx))
				ctx = ctxutil.WithValue(ctx, tt.author)
			}
			err := app.CreateChannel(ctx, tt.cmd)
			if tt.hasErr {
				assert.Error(t, err)
				return
			} else if tt.wantErr != nil {
				assert.Error(t, err)
				assert.True(t, errors.Is(err, tt.wantErr))
				return
			}
			assert.NoError(t, err)
			assert.Len(t, events, 1)
			e, ok := events[0].(stack.CreateChannelEvent)
			assert.True(t, ok)
			assert.Equal(t, vo2.Name(tt.cmd.Name), e.Name())
			assert.Equal(t, post.AuthorID(tt.author), e.AuthorID())
			assert.Greater(t, e.CreatedAt(), now)
			assert.NotEqual(t, "", e.ID())
		})
	}
}
