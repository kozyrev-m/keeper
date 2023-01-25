package httpclient

// responseError - scheme for response error.
type responseError struct {
	Error string `json:"error"`
}

// responseUser - scheme for response user.
type responseUser struct {
	ID int `json:"id"` 
	Login string `json:"login"`
}