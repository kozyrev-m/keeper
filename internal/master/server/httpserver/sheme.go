package httpserver

// requestUser - scheme for request user.
type requestUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// requestText - scheme for request text.
type requestText struct {
	Text     string `json:"text"`
	Metadata string `json:"metadata"`
}

// requestPair - scheme for request login password pair.
type requestPair struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

// requestBankCard - scheme for request bank card.
type requestBankCard struct {
	PAN       uint64 `json:"pan"` // PAN (primary account number)
	CVV       uint8  `json:"cvv"` // CVV/CVC (Card Verification Value/Code)
	ValidThru string `json:"valid_thru"`
	Name      string `json:"name"`

	Metadata string `json:"metadata"`
}
