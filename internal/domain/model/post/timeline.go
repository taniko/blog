package post

import "time"

type TimelineID string

type Timeline struct {
	id         TimelineID
	authorID   AuthorID
	visibility Visibility
	state      State
	createdAt  time.Time
	updatedAt  time.Time
}

type (
	StackID string
)

type Stack struct {
	id        StackID
	timeline  TimelineID
	text      string
	createdAt time.Time
	updatedAt time.Time
}

func CreateTimeline(id TimelineID, authorID AuthorID, visibility Visibility, state State) Timeline {
	return Timeline{
		id:         id,
		authorID:   authorID,
		visibility: visibility,
		state:      state,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}
}

func (t Timeline) CreateStack(id StackID, text string) Stack {
	return Stack{
		id:        id,
		timeline:  t.id,
		text:      text,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}
