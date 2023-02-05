package model

// Pair contains login-password pair.
type Pair struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
