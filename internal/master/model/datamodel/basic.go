package datamodel

// BasePart contains the same information for all data records.
type BasePart struct {
	ID int
	OwnerID int
	TypeID int
	Metadata string
	EncodedContent string
}

// SetID sets ID.
func (b *BasePart) SetID(id int) {
	b.ID = id
}

// GetID gets ID.
func (b *BasePart) GetID() int {
	return b.ID
}

// SetOwnerID sets OwnerID.
func (b *BasePart) SetOwnerID(ownerid int) {
	b.OwnerID = ownerid
}

// GetOwnerID gets OwnerID.
func (b *BasePart) GetOwnerID() int {
	return b.OwnerID
}

// SetTypeID sets TypeID.
func (b *BasePart) SetTypeID(typeid int) {
	b.TypeID = typeid
}

// GetTypeID gets TypeID.
func (b *BasePart) GetTypeID() int {
	return b.TypeID
}

// SetMetadata sets Metadata.
func (b *BasePart) SetMetadata(metadata string) {
	b.Metadata = metadata
}

// GetMetadata gets Metadata.
func (b *BasePart) GetMetadata() string {
	return b.Metadata
}

// SetEncodedContent sets EncodedContent.
func (b *BasePart) SetEncodedContent(content string) {
	b.Metadata = content
}

// GetEncodedContent gets EncodedContent.
func (b *BasePart) GetEncodedContent() string {
	return b.EncodedContent
}