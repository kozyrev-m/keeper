package datamodel_test

import (
	"strings"
	"testing"

	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBankCard_EncryptDecrypt(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		bc := &datamodel.BankCard{
			PAN:       10001000100010001111,
			CVV:       123,
			ValidThru: "11/23",
			Name:      "John White",
		}

		err := bc.Encrypt()
		require.NoError(t, err)

		newbc := &datamodel.BankCard{}

		err = newbc.Decrypt(bc.EncryptedContent)
		require.NoError(t, err)

		assert.Equal(t, bc.PAN, newbc.PAN)
		assert.Equal(t, bc.CVV, newbc.CVV)
		assert.Equal(t, bc.ValidThru, newbc.ValidThru)
		assert.Equal(t, bc.Name, newbc.Name)
	})

	t.Run("unsuccess", func(t *testing.T) {
		bc := &datamodel.BankCard{
			PAN:       10001000100010001111,
			CVV:       123,
			ValidThru: "11/23",
			Name:      "John White",
		}

		err := bc.Encrypt()
		require.NoError(t, err)

		withsalt := strings.Join([]string{bc.EncryptedContent, "salt"}, "")

		newbc := &datamodel.BankCard{}
		err = newbc.Decrypt(withsalt)
		require.Error(t, err)

		assert.NotEqual(t, bc.PAN, newbc.PAN)
		assert.NotEqual(t, bc.CVV, newbc.CVV)
		assert.NotEqual(t, bc.ValidThru, newbc.ValidThru)
		assert.NotEqual(t, bc.Name, newbc.Name)
	})
}
