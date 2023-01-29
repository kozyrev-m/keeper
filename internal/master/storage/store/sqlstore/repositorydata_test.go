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
		BasePart: datamodel.BasePart{
			OwnerID: 1,
			TypeID: 1,
			Metadata: "some metadata",
		},
		Value: "some text for testing",
	}

	err := s.CreateDataRecord(text)
	require.NoError(t, err)

	assert.NotEmpty(t, text.ID)
}

func TestDataRepository_FindTextsByOwner(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown()
	
	s := sqlstore.New(db)

	text := &datamodel.Text{
		BasePart: datamodel.BasePart{
			OwnerID: 1,
			TypeID: 1,
			Metadata: "some metadata",
		},
		Value: "some text for testing",
	}

	err := s.CreateDataRecord(text)
	require.NoError(t, err)

	assert.NotEmpty(t, text.ID)
	
	texts, err := s.FindTextsByOwner(1)
	require.NoError(t, err)

	assert.Equal(t, text.Value, texts[0].Value)
}