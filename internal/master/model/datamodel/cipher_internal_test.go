package datamodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testpassword = "keyfortest"
)

func TestCipher_Encrypt(t *testing.T) {
	testvalue := "some text for testing"
	
	enc, err := encrypt(testpassword, testvalue)
	require.NoError(t, err)

	result, err := decrypt(testpassword, enc)
	require.NoError(t, err)

	assert.Equal(t, testvalue, result)
}