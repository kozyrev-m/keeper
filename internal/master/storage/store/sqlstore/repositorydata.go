package sqlstore

import (
	"database/sql"
	"log"

	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/kozyrev-m/keeper/internal/master/storage/store"
)

const (
	limit = 10
)

// Content ...
type Content interface {
	Encrypt() error
	SetID(int)
	GetOwnerID() int
	GetTypeID() int
	GetMetadata() string
	GetEncryptedContent() string
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
		c.GetEncryptedContent(),
	).Scan(&id); err != nil {
		return err
	}

	c.SetID(id)

	return nil
}

// FindTextsByOwner gets all texts.
func (s *Store) FindTextsByOwner(ownerid int) ([]datamodel.Text, error) {
	baseParts, err := s.findRecords(ownerid, 1)
	if err != nil {
		return nil, err
	}

	texts := make([]datamodel.Text, 0, limit)
	for _, base := range baseParts {
		text := datamodel.Text{
			BasePart: base,
		}

		if err := text.Decrypt(text.EncryptedContent); err != nil {
			return nil, err
		}

		texts = append(texts, text)
	}

	return texts, nil
}

// FindRecords gets data records by owner id and data type.
func (s *Store) findRecords(ownerID int, typeID int) ([]datamodel.BasePart, error) {
	baseParts := make([]datamodel.BasePart, 0, limit)

	rows, err := s.db.Query(
		"SELECT id, owner_id, type_id, metadata, content FROM private_data WHERE owner_id = $1 AND type_id = $2",
		ownerID,
		typeID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	defer func () {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	} ()

	for rows.Next() {
		b := datamodel.BasePart{}
		
		if err := rows.Scan(&b.ID, &b.OwnerID, &b.TypeID, &b.Metadata, &b.EncryptedContent); err != nil {
			return nil, err
		}
		baseParts = append(baseParts, b)
	}

	return baseParts, nil
}