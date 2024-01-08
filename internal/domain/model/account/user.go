package account

type (
	ID string
)

type Account struct {
	id ID
}

func NewAccount(id ID) Account {
	return Account{
		id: id,
	}
}
