package account

type (
	Name string
	Bio  string
)

type Profile struct {
	id   ID
	name Name
	bio  Bio
}

func NewProfile(id ID, name Name, bio Bio) Profile {
	return Profile{
		id:   id,
		name: name,
		bio:  bio,
	}
}
