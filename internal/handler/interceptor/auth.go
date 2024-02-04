package interceptor

import (
	"context"
	"errors"
	"strings"

	"connectrpc.com/connect"
	application "github.com/taniko/blog/internal/application/auth"
	"github.com/taniko/blog/internal/utils/ctxutil"
)

const tokenHeader = "authorization"

type AuthInterceptor interface {
	UnaryInterceptorFunc() connect.UnaryInterceptorFunc
}

type authInterceptor struct {
	s application.Service
}

func NewAuthInterceptor(s application.Service) AuthInterceptor {
	return &authInterceptor{s: s}
}

func (a authInterceptor) UnaryInterceptorFunc() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			token, err := a.extractToken(req.Header().Get(tokenHeader))
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid token"),
				)
			}
			uid, err := a.s.Verify(ctx, token).Get()
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("invalid token"),
				)
			}
			ctx = ctxutil.WithValue(ctx, uid)
			return next(ctx, req)
		}
	}
}

// Bearerのトークン部分を抽出する
func (a authInterceptor) extractToken(bearerToken string) (string, error) {
	parts := strings.Split(bearerToken, "Bearer ")
	if len(parts) != 2 {
		return "", errors.New("invalid token")
	}
	return parts[1], nil
}
