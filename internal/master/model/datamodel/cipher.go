package datamodel

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

// initCipher prepares cipher to encrypt/decrypt.
func initCipher(password string) (cipher.AEAD, []byte, error) {
	key := sha256.Sum256([]byte(password))

    aesblock, err := aes.NewCipher(key[:])
    if err != nil {
        return nil, nil, err
    }
    aesgcm, err := cipher.NewGCM(aesblock)
    if err != nil {
		return nil, nil, err
    }

	// creating an initialization vector
    nonce := key[len(key)-aesgcm.NonceSize():]

	return aesgcm, nonce, nil
}

// encrypt encrypts some string.
func encrypt(password string, value string) (string, error) {
	aesgcm, nonce, err := initCipher(password)
	if err != nil {
		return "", err
	}

    dst := aesgcm.Seal(nil, nonce, []byte(value), nil) // encrypting

	return hex.EncodeToString(dst), nil
}

// decrypt decrypts encrypted string.
func decrypt(password string, enc string) (string, error) {
	aesgcm, nonce, err := initCipher(password)
	if err != nil {
		return "", err
	}

	encrypted, err := hex.DecodeString(enc)
    if err != nil {
    	return "", err
    }
	
    decrypted, err := aesgcm.Open(nil, nonce, encrypted, nil) // decoding
    if err != nil {
		return "", err
	}

	return string(decrypted), nil
}