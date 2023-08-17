package papershamir

import (
	"testing"

	"github.com/matryer/is"
)

func TestEncrypt(t *testing.T) {

	t.Run("can encrypt and decrypt message", func(t *testing.T) {
		is := is.New(t)

		encrypted, err := encrypt.encrypt("shh", []byte("Alphabet"))
		is.NoErr(err)

		decrypted, err := encrypt.decrypt("shh", encrypted)
		is.NoErr(err)
		is.Equal(string(decrypted), "Alphabet")
	})

	t.Run("cannot decrypt with wrong passphrase", func(t *testing.T) {
		is := is.New(t)

		encrypted, err := encrypt.encrypt("shh", []byte("Alphabet"))
		is.NoErr(err)

		_, err = encrypt.decrypt("sh", encrypted)
		is.Equal(err.Error(), "cipher: message authentication failed")
	})

	t.Run("cannot decrypt with wrong salt", func(t *testing.T) {
		is := is.New(t)

		encrypted, err := encrypt.encrypt("shh", []byte("Alphabet"))
		is.NoErr(err)
		encrypted[1] = encrypted[1] + 1

		_, err = encrypt.decrypt("shh", encrypted)
		is.Equal(err.Error(), "cipher: message authentication failed")
	})

	t.Run("cannot decrypt with wrong nonce", func(t *testing.T) {
		is := is.New(t)

		encrypted, err := encrypt.encrypt("shh", []byte("Alphabet"))
		is.NoErr(err)
		encrypted[9] = encrypted[9] + 1

		_, err = encrypt.decrypt("shh", encrypted)
		is.Equal(err.Error(), "cipher: message authentication failed")
	})

	t.Run("cannot decrypt with wrong data", func(t *testing.T) {
		is := is.New(t)

		encrypted, err := encrypt.encrypt("shh", []byte("Alphabet"))
		is.NoErr(err)
		encrypted[len(encrypted)-1] = encrypted[len(encrypted)-1] + 1

		_, err = encrypt.decrypt("shh", encrypted)
		is.Equal(err.Error(), "cipher: message authentication failed")
	})

	t.Run("cannot decrypt with small data", func(t *testing.T) {
		is := is.New(t)

		_, err := encrypt.decrypt("sh", []byte{1})
		is.Equal(err.Error(), "crypto data invalid")

		_, err = encrypt.decrypt("sh", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9})
		is.Equal(err.Error(), "crypto data invalid")
	})
}
