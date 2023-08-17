package papershamir

import (
	"testing"

	"github.com/bradenrayhorn/papershamir/internal/testutils"
	"github.com/matryer/is"
)

func TestSplit(t *testing.T) {

	t.Run("splits small secret into proper format", func(t *testing.T) {
		is := is.New(t)

		output, err := Split([]byte{'A'}, 2, 2, "")
		is.NoErr(err)
		is.Equal(len(output), 2)

		r := testutils.CreateShareRegexp(1, 4, 4)
		for _, share := range output {
			is.True(r.MatchString(share))
		}
	})

	t.Run("wraps larger secret", func(t *testing.T) {
		is := is.New(t)

		output, err := Split([]byte("12345678901234567890"), 2, 2, "")
		is.NoErr(err)
		is.Equal(len(output), 2)

		r := testutils.CreateShareRegexp(2, 40, 2)
		for _, share := range output {
			is.True(r.MatchString(share))
		}
	})

	t.Run("fails to split if configured incorrectly", func(t *testing.T) {
		is := is.New(t)

		_, err := Split([]byte("1"), 1, 1, "")
		is.Equal(err.Error(), "threshold must be at least 2")
	})

	t.Run("fails to split with empty secret", func(t *testing.T) {
		is := is.New(t)

		_, err := Split([]byte{}, 2, 2, "")
		is.Equal(err.Error(), "cannot split an empty secret")
	})

	t.Run("encrypts if passed key", func(t *testing.T) {
		is := is.New(t)

		shares, err := Split([]byte{'A'}, 2, 2, "key")
		is.NoErr(err)

		byteShares := make([][]byte, len(shares))
		for i, v := range shares {
			byteShares[i] = []byte(v)
		}

		result, err := Combine(byteShares, "key")
		is.NoErr(err)
		is.Equal(result, []byte{'A'})
	})
}
