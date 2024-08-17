package papershamir

import (
	"encoding/base64"
	"testing"

	"github.com/bradenrayhorn/papershamir/internal/testutils"
	"github.com/matryer/is"
)

func TestSplitHexr(t *testing.T) {

	t.Run("splits small secret into proper format", func(t *testing.T) {
		is := is.New(t)

		output, err := SplitHexr([]byte{'A'}, 2, 2, "")
		is.NoErr(err)
		is.Equal(len(output), 2)

		r := testutils.CreateShareRegexp(1, 6, 6)
		for _, share := range output {
			is.True(r.MatchString(share))
		}
	})

	t.Run("wraps larger secret", func(t *testing.T) {
		is := is.New(t)

		output, err := SplitHexr([]byte("1234567890123456789"), 2, 2, "")
		is.NoErr(err)
		is.Equal(len(output), 2)

		r := testutils.CreateShareRegexp(2, 40, 2)
		for _, share := range output {
			is.True(r.MatchString(share))
		}
	})

	t.Run("fails to split if configured incorrectly", func(t *testing.T) {
		is := is.New(t)

		_, err := SplitHexr([]byte("1"), 1, 1, "")
		is.Equal(err.Error(), "threshold must be at least 2")
	})

	t.Run("fails to split with empty secret", func(t *testing.T) {
		is := is.New(t)

		_, err := SplitHexr([]byte{}, 2, 2, "")
		is.Equal(err.Error(), "cannot split an empty secret")
	})

	t.Run("encrypts if passed key", func(t *testing.T) {
		is := is.New(t)

		shares, err := SplitHexr([]byte{'A'}, 2, 2, "key")
		is.NoErr(err)

		byteShares := make([][]byte, len(shares))
		for i, v := range shares {
			byteShares[i] = []byte(v)
		}

		result, err := CombineHexr(byteShares, "key")
		is.NoErr(err)
		is.Equal(result, []byte{'A'})
	})

	t.Run("compresses a big secret", func(t *testing.T) {
		is := is.New(t)

		secret := "AA AA AA AA AA AA AA AA AA AA AA AA AA AA AA AA"

		output, err := SplitHexr([]byte(secret), 2, 2, "")
		is.NoErr(err)
		is.Equal(len(output), 2)

		byteShares := make([][]byte, len(output))
		r := testutils.CreateShareRegexp(2, 40, 10)
		for i, share := range output {
			is.True(r.MatchString(share))
			byteShares[i] = []byte(share)
		}

		result, err := CombineHexr(byteShares, "")
		is.NoErr(err)
		is.Equal(result, []byte(secret))
	})
}

func TestSplitQR(t *testing.T) {

	t.Run("splits small secret into proper format", func(t *testing.T) {
		is := is.New(t)

		output, err := SplitQR([]byte{'A'}, 2, 2, "")
		is.NoErr(err)
		is.Equal(len(output), 2)

		for _, share := range output {
			_, err := base64.RawStdEncoding.Strict().DecodeString(share)
			is.NoErr(err)
		}
	})

	t.Run("fails to split if configured incorrectly", func(t *testing.T) {
		is := is.New(t)

		_, err := SplitHexr([]byte("1"), 1, 1, "")
		is.Equal(err.Error(), "threshold must be at least 2")
	})

	t.Run("fails to split with empty secret", func(t *testing.T) {
		is := is.New(t)

		_, err := SplitHexr([]byte{}, 2, 2, "")
		is.Equal(err.Error(), "cannot split an empty secret")
	})

	t.Run("encrypts if passed key", func(t *testing.T) {
		is := is.New(t)

		shares, err := SplitQR([]byte{'A'}, 2, 2, "key")
		is.NoErr(err)

		byteShares := make([][]byte, len(shares))
		for i, v := range shares {
			byteShares[i] = []byte(v)
		}

		result, err := CombineQR(byteShares, "key")
		is.NoErr(err)
		is.Equal(result, []byte{'A'})
	})

	t.Run("compresses a big secret", func(t *testing.T) {
		is := is.New(t)

		secret := "AA AA AA AA AA AA AA AA AA AA AA AA AA AA AA AA"

		output, err := SplitQR([]byte(secret), 2, 2, "")
		is.NoErr(err)
		is.Equal(len(output), 2)

		byteShares := make([][]byte, len(output))
		for i, share := range output {
			_, err := base64.RawStdEncoding.Strict().DecodeString(share)
			is.NoErr(err)
			byteShares[i] = []byte(share)
		}

		result, err := CombineQR(byteShares, "")
		is.NoErr(err)
		is.Equal(result, []byte(secret))
	})
}
