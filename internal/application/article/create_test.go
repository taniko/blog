package article

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/blog/internal/application/article/command"
	"github.com/taniko/blog/internal/domain/model/account"
	articlemock "github.com/taniko/blog/internal/domain/repository/mock"
	"github.com/taniko/blog/internal/utils/ctxutil"
	"go.uber.org/mock/gomock"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := articlemock.NewMockArticle(ctrl)
	repo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	tests := []struct {
		ctx    context.Context
		name   string
		cmd    command.Create
		hasErr bool
	}{
		{
			name: "success",
			ctx:  ctxutil.WithValue(context.Background(), account.ID(uuid.NewString())),
			cmd: command.Create{
				Title: "title",
				Body:  "body",
			},
		}, {
			name: "no account id",
			ctx:  context.Background(),
			cmd: command.Create{
				Title: "title",
				Body:  "body",
			},
			hasErr: true,
		},
	}

	s := New(repo)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.Create(tt.ctx, tt.cmd)
			if tt.hasErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
