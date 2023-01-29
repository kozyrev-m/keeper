package datamodel

// BasePart contains similar information for all data records.
type BasePart struct {
	ID int `json:"id"`
	OwnerID int `json:"owner_id"`
	TypeID int `json:"type_id"`
	Metadata string `json:"metadata"`
	EncryptedContent string `json:"-"`
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

// SetEncodedContent sets EncryptedContent.
func (b *BasePart) SetEncryptedContent(content string) {
	b.EncryptedContent = content
}

// GetEncodedContent gets EncryptedContent.
func (b *BasePart) GetEncryptedContent() string {
	return b.EncryptedContent
}