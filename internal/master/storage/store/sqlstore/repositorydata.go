package sqlstore


type Content interface {
	Encrypt() error
	SetID(int)
	GetOwnerID() int
	GetTypeID() int
	GetMetadata() string
	GetEncodedContent() string
}

// CreateDataRecord creates record with content.
func (s *Store) CreateDataRecord(c Content) error {
	
	if err := c.Encrypt(); err != nil {
		return err
	}
	
	var id int
	if err := s.db.QueryRow(
		"INSERT INTO private_data (owner_id, type_id, metadata, content) VALUES ($1, $2, $3, $4) RETURNING id",
		c.GetOwnerID(),
		c.GetTypeID(),
		c.GetMetadata(),
		c.GetEncodedContent(),
	).Scan(&id); err != nil {
		return err
	}

	c.SetID(id)

	return nil
}
