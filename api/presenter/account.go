package presenter

// Account represets an user account
type Account struct {
	ID                   int     `json:"id"`
	DocumentNumber       string  `json:"document_number"`
	AvailableCreditLimit float64 `json:"available_credit_limit"`
}
