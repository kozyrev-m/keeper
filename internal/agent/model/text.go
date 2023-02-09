package model

// Text contains text.
type Text struct {
	Metadata string `metadata:"metadata"`
	Value    string `json:"text"`
}
