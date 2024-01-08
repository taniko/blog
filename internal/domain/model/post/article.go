package post

import "time"

type Article struct {
	id         ID
	author     AuthorID
	title      Title
	body       Body
	createdAt  time.Time
	updatedAt  time.Time
	state      State
	visibility Visibility
}
