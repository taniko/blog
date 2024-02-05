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

func CreateArticle(id ID, author AuthorID, title Title, body Body, state State, visibility Visibility) Article {
	return Article{
		id:         id,
		author:     author,
		title:      title,
		body:       body,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
		state:      state,
		visibility: visibility,
	}
}
