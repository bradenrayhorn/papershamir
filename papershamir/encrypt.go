package papershamir

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

type encryptImpl struct {
}

var encrypt = encryptImpl{}

func (e *encryptImpl) encrypt(passphrase string, data []byte) ([]byte, error) {
	salt, err := createSalt()
	if err != nil {
		return nil, err
	}

	key, err := createKey(passphrase, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	encrypted := gcm.Seal(nil, nonce, data, nil)
	return append(append(salt, nonce...), encrypted...), nil
}

func (e *encryptImpl) decrypt(passphrase string, encrypted []byte) ([]byte, error) {
	if len(encrypted) < 9 {
		return nil, fmt.Errorf("crypto data invalid")
	}

	salt := encrypted[:8]
	encrypted = encrypted[8:]

	key, err := createKey(passphrase, salt)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < gcm.NonceSize() {
		return nil, fmt.Errorf("crypto data invalid")
	}
	nonce := encrypted[:gcm.NonceSize()]
	encrypted = encrypted[gcm.NonceSize():]

	return gcm.Open(nil, nonce, encrypted, nil)
}

func createKey(passphrase string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(passphrase), salt, 32768, 8, 1, 32)
}

func createSalt() ([]byte, error) {
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
