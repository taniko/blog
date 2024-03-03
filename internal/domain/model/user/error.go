package user

import "github.com/cockroachdb/errors"

var ErrUnauthenticated = errors.New("unauthenticated")
