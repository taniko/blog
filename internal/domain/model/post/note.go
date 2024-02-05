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

func CreateNote(id ID, title Title, body Body, state State, visibility Visibility) Note {
	return Note{
		id:         id,
		title:      title,
		body:       body,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
		state:      state,
		visibility: visibility,
	}
}
