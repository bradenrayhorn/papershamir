package papershamir

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

type hexrImpl struct {
	substitutions map[byte]byte
}

var hexr = hexrImpl{
	substitutions: map[byte]byte{
		'B': 'W',
		'C': 'X',
		'D': 'H',
		'F': 'K',
		'0': 'N',
	},
}

func (h hexrImpl) encode(secret []byte) []byte {
	// encode secret to hex
	encoded := make([]byte, hex.EncodedLen(len(secret)))
	hex.Encode(encoded, secret)

	// turn hex into all uppercase
	encoded = bytes.ToUpper(encoded)

	// perform substitutions
	for i, c := range encoded {
		for original, replacement := range h.substitutions {
			if c == original {
				encoded[i] = replacement
			}
		}
	}

	return encoded
}

func (h hexrImpl) decode(encoded []byte) ([]byte, error) {
	// perform substitutions
	for i, c := range encoded {
		for replacement, original := range h.substitutions {
			if c == original {
				encoded[i] = replacement
			}
		}
	}

	// turn secret into lowercase
	encoded = bytes.ToLower(encoded)

	// decode hex
	secret := make([]byte, hex.DecodedLen(len(encoded)))
	p, err := hex.Decode(secret, encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode at character %d of \"%s\": %w", p, encoded, err)
	}

	return secret, nil
}
