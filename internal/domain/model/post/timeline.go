package post

import "time"

type TimelineID string

type Timeline struct {
	id         TimelineID
	authorID   AuthorID
	stacks     []Stack
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
