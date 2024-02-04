package ctxutil

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	type (
		ID     string
		Number int
		Name   string
	)

	ctx := context.Background()
	ctx = WithValue(ctx, ID("id"))
	ctx = WithValue(ctx, Number(1))
	ctx = WithValue(ctx, Name("name"))

	id, ok := Value[ID](ctx)
	assert.True(t, ok)
	assert.Equal(t, ID("id"), id)

	num, ok := Value[Number](ctx)
	assert.True(t, ok)
	assert.Equal(t, Number(1), num)

	name, ok := Value[Name](ctx)
	assert.True(t, ok)
	assert.Equal(t, Name("name"), name)
}
