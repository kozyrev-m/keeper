// Package datamodel contains model of data record.
package datamodel

// Content is content iterface.
type Content interface {
	Encrypt() (string, error)
	Decrypt(string) error
}

// DataRecord contains data record.
type DataRecord struct {
	ID int
	TypeID int
	OwnerID int

	Content Content // Main part of the record

	EncodedContent string

	Metadata string
}
