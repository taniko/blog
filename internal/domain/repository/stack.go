//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"context"

	"github.com/samber/mo"
	event "github.com/taniko/blog/internal/domain/event/stack"
	"github.com/taniko/blog/internal/domain/model/stack"
	"github.com/taniko/blog/internal/domain/model/stack/vo"
)

type Stack interface {
	// Save イベントを保存する
	Save(ctx context.Context, events []event.Event) error

	// FindChannel チャンネルを取得する
	FindChannel(ctx context.Context, id vo.ChannelID) mo.Result[stack.Channel]

	// GetMessages メッセージを取得する
	GetMessages(ctx context.Context, channelID vo.ChannelID) mo.Result[[]stack.Message]
}
