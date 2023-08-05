package papershamir

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestCombine(t *testing.T) {

	t.Run("can combine shares", func(t *testing.T) {
		is := is.New(t)

		shares := [][]byte{
			[]byte("99 4E XNKN413A\n      XNKN413A"),
			[]byte("97 K7 E3A87326\n      E3A87326"),
		}

		result, err := Combine(shares)
		is.NoErr(err)
		is.Equal(string(result), "x")
	})

	t.Run("fails if only one share provided", func(t *testing.T) {
		is := is.New(t)

		shares := [][]byte{
			[]byte("99 4E XNKN413A\n      XNKN413A"),
		}

		_, err := Combine(shares)
		is.Equal(err.Error(), "less than two parts cannot be used to reconstruct the secret")
	})

	t.Run("fails with incomplete share", func(t *testing.T) {
		is := is.New(t)

		shares := [][]byte{
			[]byte("99 4E XNKN413A"),
		}

		_, err := Combine(shares)
		is.Equal(err.Error(), "share requires at least two lines. at share 1")
	})

	t.Run("can detect line checksum failure", func(t *testing.T) {
		is := is.New(t)

		shares := [][]byte{
			[]byte("99 4E XNKN414A\n      XNKN413A"),
			[]byte("97 K7 E3A87326\n      E3A87326"),
		}

		_, err := Combine(shares)
		is.Equal(err.Error(), "checksum failed. at share 1 line 1")
	})

	t.Run("can detect share checksum failure", func(t *testing.T) {
		is := is.New(t)

		shares := [][]byte{
			[]byte("99 4E XNKN413A\n      XNKN413A"),
			[]byte("97 K7 E3A87326\n      E3A87336"),
		}

		_, err := Combine(shares)
		is.Equal(err.Error(), "share checksum failed. at share 2")
	})

	t.Run("can handle share checksum decode error", func(t *testing.T) {
		is := is.New(t)

		shares := [][]byte{
			[]byte("99 4E XNKN413A\n      -NKN413A"),
			[]byte("97 K7 E3A87326\n      E3A87336"),
		}

		_, err := Combine(shares)
		is.True(strings.Contains(err.Error(), "failed to decode at character 0"))
	})

	t.Run("can handle line checksum decode error", func(t *testing.T) {
		is := is.New(t)

		shares := [][]byte{
			[]byte("99 4E -NKN413A\n      XNKN413A"),
			[]byte("97 K7 E3A87326\n      E3A87336"),
		}

		_, err := Combine(shares)
		is.True(strings.Contains(err.Error(), "failed to decode at character 0"))
	})

}
