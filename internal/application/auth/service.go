package auth

import (
	"context"

	"firebase.google.com/go/v4/auth"
	"github.com/cockroachdb/errors"
	"github.com/samber/mo"
	"github.com/taniko/blog/internal/domain/model/user/vo"
)

type Service interface {
	Verify(ctx context.Context, token string) mo.Result[vo.UserID]
}

type service struct {
	c *auth.Client
}

func NewService(c *auth.Client) Service {
	return &service{c: c}
}

func (s service) Verify(ctx context.Context, in string) mo.Result[vo.UserID] {
	token, err := s.c.VerifyIDToken(ctx, in)
	if err != nil {
		return mo.Err[vo.UserID](errors.Wrap(err, "failed to verify token"))
	}
	return mo.Ok(vo.UserID(token.UID))
}
