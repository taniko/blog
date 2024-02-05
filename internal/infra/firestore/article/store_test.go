//go:build firebase

package article

import (
	"context"
	"testing"

	firebase "firebase.google.com/go/v4"
	"github.com/stretchr/testify/assert"
	"github.com/taniko/blog/internal/domain/event/article"
)

func TestArticle_Store(t *testing.T) {
	ctx := context.Background()
	conf := &firebase.Config{}
	app, err := firebase.NewApp(ctx, conf)
	assert.NoError(t, err)
	client, err := app.Firestore(ctx)
	assert.NoError(t, err)
	defer client.Close()

	type test struct {
		name   string
		event  article.Event
		hasErr bool
	}

	var tests []test

	tests = append(tests, test{
		name:  "success",
		event: article.NewCreateEvent("author", "title", "body"),
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewArticleStore(client)
			err := a.Store(ctx, tt.event)
			if tt.hasErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
