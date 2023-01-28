package sqlstore

import (
	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
)

// CreateDataRecord creates data record.
func (s *Store) CreateDataRecord(d *datamodel.DataRecord) error {
	if err := d.BeforeCreate(); err != nil {
		return err
	}

	return s.db.QueryRow(
		"INSERT INTO private_data (owner_id, type_id, metadata, content) VALUES ($1, $2, $3, $4) RETURNING id",
		d.OwnerID,
		d.TypeID,
		d.Metadata,
		d.EncodedContent,
	).Scan(&d.ID)
}
