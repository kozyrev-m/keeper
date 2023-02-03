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
