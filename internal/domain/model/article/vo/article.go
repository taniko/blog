package vo

type (
	ID       string
	Title    string
	Body     string
	AuthorID string
)

func (i ID) String() string {
	return string(i)
}

func (t Title) String() string {
	return string(t)
}

func (b Body) String() string {
	return string(b)
}

func (a AuthorID) String() string {
	return string(a)
}
