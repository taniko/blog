package dto

import (
	"time"

	"github.com/taniko/blog/internal/domain/event"
)

type EventHeader struct {
	ID      string    `firestore:"event_id"`
	Time    time.Time `firestore:"created_at"`
	Version int       `firestore:"version"`
}

func NewEventHeader(h event.Header) EventHeader {
	return EventHeader{
		ID:      h.GetID().String(),
		Time:    h.GetTime(),
		Version: h.GetVersion().Int(),
	}
}
