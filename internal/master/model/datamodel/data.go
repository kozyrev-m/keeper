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

// New creates data record with content implementation.
func New(typeid int) *DataRecord {
	r := &DataRecord{
		TypeID: typeid,
	}
	if typeid == 1 { // text has typeid = 1
		r.Content = &Text{}
	}

	return r
}

// BeforeCreate prepares data record to insert to store.
func (d *DataRecord) BeforeCreate() error {
	enc, err := d.Content.Encrypt()
	if err != nil {
		return err
	}

	d.EncodedContent = enc

	return nil
}

// AfterReceive prepares data record after receive from store.
func (d *DataRecord) AfterReceive() error {
	return d.Content.Decrypt(d.EncodedContent)
}
