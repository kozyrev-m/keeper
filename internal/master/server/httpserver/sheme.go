package httpserver

// request - scheme for request user.
type requestUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}