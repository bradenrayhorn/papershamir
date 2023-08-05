package papershamir

import (
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestHexr(t *testing.T) {
	t.Run("can encode to hexr", func(t *testing.T) {
		is := is.New(t)

		res := hexr.encode([]byte{'j', 'k', 'l', 'm', 'n', 'o', 'p'})
		is.Equal("6A6W6X6H6E6K7N", string(res))
	})

	t.Run("can decode to hexr", func(t *testing.T) {
		is := is.New(t)

		res, err := hexr.decode([]byte("6A6W6X6H6E6K7N"))
		is.NoErr(err)
		is.Equal([]byte{'j', 'k', 'l', 'm', 'n', 'o', 'p'}, res)
	})

	t.Run("handles invalid characters", func(t *testing.T) {
		is := is.New(t)

		_, err := hexr.decode([]byte("~!!-190  13 -"))
		is.True(strings.Contains(err.Error(), "failed to decode"))
	})
}
