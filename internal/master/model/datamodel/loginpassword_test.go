package datamodel_test

import (
	"strings"
	"testing"

	"github.com/kozyrev-m/keeper/internal/master/model/datamodel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginPassword_EncryptDecrypt(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		lp := &datamodel.LoginPassword{
			Login:    "someuser",
			Password: "password",
		}

		err := lp.Encrypt()
		require.NoError(t, err)

		newlp := &datamodel.LoginPassword{}
		err = newlp.Decrypt(lp.EncryptedContent)
		require.NoError(t, err)

		assert.Equal(t, lp.Login, newlp.Login)
		assert.Equal(t, lp.Password, newlp.Password)
	})

	t.Run("unsuccess", func(t *testing.T) {
		lp := &datamodel.LoginPassword{
			Login:    "someuser",
			Password: "password",
		}

		err := lp.Encrypt()
		require.NoError(t, err)

		withsalt := strings.Join([]string{lp.EncryptedContent, "salt"}, "")

		newlp := &datamodel.LoginPassword{}
		err = newlp.Decrypt(withsalt)
		require.Error(t, err)

		assert.NotEqual(t, lp.Login, newlp.Login)
		assert.NotEqual(t, lp.Password, newlp.Password)
	})
}
