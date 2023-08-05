package papershamir

import (
	"testing"

	"github.com/matryer/is"
)

func TestChecksum(t *testing.T) {
	t.Run("can create and verify checksum", func(t *testing.T) {
		is := is.New(t)

		c := checksum.create("ABC")
		is.Equal(true, checksum.verify("ABC", c))
	})

	t.Run("fails if passed different secret", func(t *testing.T) {
		is := is.New(t)

		c := checksum.create("ABC")
		is.Equal(false, checksum.verify("ABCD", c))
	})

	t.Run("fails if passed different checksum", func(t *testing.T) {
		is := is.New(t)

		is.Equal(false, checksum.verify("ABC", []byte{'1'}))
	})

	t.Run("fails if passed invalid checksum", func(t *testing.T) {
		is := is.New(t)

		is.Equal(false, checksum.verify("ABC", []byte{}))
	})
}
