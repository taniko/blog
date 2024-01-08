package post

import "time"

type Note struct {
	id         ID
	title      Title
	body       Body
	createdAt  time.Time
	updatedAt  time.Time
	state      State
	visibility Visibility
}
