package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/kozyrev-m/keeper/internal/master/storage/store"
	"github.com/kozyrev-m/keeper/internal/master/storage/store/filestorage"
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
	baseParts, err := s.findRecords(ownerid, datamodel.TypeText)
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

// FindPairsByOwner gets all login password pairs.
func (s *Store) FindPairsByOwner(ownerid int) ([]datamodel.LoginPassword, error) {
	baseParts, err := s.findRecords(ownerid, datamodel.TypePair)
	if err != nil {
		return nil, err
	}

	pairs := make([]datamodel.LoginPassword, 0, limit)
	for _, base := range baseParts {
		loginPassword := datamodel.LoginPassword{
			BasePart: base,
		}

		if err := loginPassword.Decrypt(loginPassword.EncryptedContent); err != nil {
			return nil, err
		}

		pairs = append(pairs, loginPassword)
	}

	return pairs, nil
}

// FindBankCardsByOwner gets all bank cards.
func (s *Store) FindBankCardsByOwner(ownerid int) ([]datamodel.BankCard, error) {
	baseParts, err := s.findRecords(ownerid, datamodel.TypeBank)
	if err != nil {
		return nil, err
	}

	bankcards := make([]datamodel.BankCard, 0, limit)
	for _, base := range baseParts {
		bankCard := datamodel.BankCard {
			BasePart: base,
		}

		if err := bankCard.Decrypt(bankCard.EncryptedContent); err != nil {
			return nil, err
		}

		bankcards = append(bankcards, bankCard)
	}

	return bankcards, nil
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

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		b := datamodel.BasePart{}

		if err := rows.Scan(&b.ID, &b.OwnerID, &b.TypeID, &b.Metadata, &b.EncryptedContent); err != nil {
			return nil, err
		}
		baseParts = append(baseParts, b)
	}

	return baseParts, nil
}

// CreateFile creates file on disk and file record on db.
func (s *Store) CreateFile(ownerID int, metadata string, filename string, file multipart.File) error {
	filepath := fmt.Sprintf("%s/%d/%s", filestorage.Dir, ownerID, filename)

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
	}()

	if _, err := tx.Exec(
		"INSERT INTO files (owner_id, metadata, filepath) VALUES ($1, $2, $3)",
		ownerID,
		metadata,
		filepath,
	); err != nil {
		return err
	}

	// create file on disk
	if err := filestorage.CreateFile(ownerID, filename, file); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		if errIn := filestorage.DeleteFile(filepath); errIn != nil {
			return errIn
		}

		return err
	}

	return nil
}

// GetFileList gets file list.
func (s *Store) GetFileList(ownerID int) ([]datamodel.File, error) {
	fileList := make([]datamodel.File, 0)
	rows, err := s.db.Query(
		"SELECT id, metadata, filepath FROM files WHERE owner_id = $1",
		ownerID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	for rows.Next() {
		f := datamodel.File{}

		if err := rows.Scan(&f.ID, &f.Metadata, &f.Filepath); err != nil {
			return nil, err
		}
		
		f.Name()

		fileList = append(fileList, f)
	}

	return fileList, nil
}