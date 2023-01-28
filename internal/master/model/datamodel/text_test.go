package datamodel_test

import (
	"strings"
	"testing"

	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestText_EncryptDecrypt(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		text := &datamodel.Text{
			Value: "some text for test",
		}
	
		enc, err := text.Encrypt()
		require.NoError(t, err)
	
		newText := &datamodel.Text{}
		newText.Decrypt(enc)
	
		assert.Equal(t, text.Value, newText.Value)
	})

	t.Run("unsuccess", func(t *testing.T) {
		text := &datamodel.Text{
			Value: "some text for test",
		}
	
		enc, err := text.Encrypt()
		require.NoError(t, err)
	
		withsalt := strings.Join([]string{enc, "salt"}, "")

		newText := &datamodel.Text{}
		newText.Decrypt(withsalt)
	
		assert.NotEqual(t, text.Value, newText.Value)
	})
}