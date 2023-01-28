package datamodel

type Data struct {
	ID int
	TypeID int
	OwnerID int

	Content Content
	EncodedContent string

	Metadata string
}
