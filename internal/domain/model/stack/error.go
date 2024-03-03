package stack

import "github.com/cockroachdb/errors"

var (
	// ErrEmptyName nameが空の場合のエラー
	ErrEmptyName = errors.New("name is empty")

	// ErrEmptyDescription descriptionが空の場合のエラー
	ErrEmptyDescription = errors.New("description is empty")
)
