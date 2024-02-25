package vo

type UserID string

func (id UserID) String() string {
	return string(id)
}
