package sqlstore_test

import (
	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/kozyrev-m/keeper/internal/master/storage/store/sqlstore"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataRepository_CreateDataRecord(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown()
	
	s := sqlstore.New(db)

	text := &datamodel.Text{
		Value: "some text for testing",
	}

	dr := &datamodel.DataRecord{
		TypeID: 1,
		OwnerID: 1,
		Metadata: "some metadata",
		Content: text,
	}

	err := s.CreateDataRecord(dr)
	require.NoError(t, err)

	assert.NoError(t, err)
	assert.NotNil(t, dr.ID)
}
