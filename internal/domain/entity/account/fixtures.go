package account

func NewFixtureAccount() *Account {
	return &Account{
		ID:                   1,
		DocumentNumber:       "12345678900",
		AvailableCreditLimit: 1000,
	}
}
