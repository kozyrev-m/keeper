package model

// Pair contains login-password pair.
type Pair struct {
	Metadata string `metadata:"metadata"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
