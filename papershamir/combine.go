package papershamir

import (
	"fmt"
	"strings"

	"github.com/hashicorp/vault/shamir"
)

func Combine(shares [][]byte) ([]byte, error) {
	decodedShares := make([][]byte, len(shares))

	for i, share := range shares {
		lines := strings.Split(string(share), "\n")

		if len(lines) < 2 {
			return nil, fmt.Errorf("share requires at least two lines. at share %d", i+1)
		}

		secret := ""
		for j, line := range lines[:len(lines)-1] {
			tokens := strings.Split(line, " ")
			foundChecksum := tokens[len(tokens)-1]
			foundSecret := strings.Join(tokens[:len(tokens)-1], "")
			secret += foundSecret

			// validate checksum for line
			decodedChecksum, err := hexr.decode([]byte(foundChecksum))
			if err != nil {
				return nil, err
			}
			if !checksum.verify(foundSecret, decodedChecksum) {
				return nil, fmt.Errorf("checksum failed. at share %d line %d", i+1, j+1)
			}
		}

		// validate share checksum
		shareChecksum := strings.TrimSpace(lines[len(lines)-1])
		decodedChecksum, err := hexr.decode([]byte(shareChecksum))
		if err != nil {
			return nil, err
		}
		if !checksum.verify(secret, decodedChecksum) {
			return nil, fmt.Errorf("share checksum failed. at share %d", i+1)
		}

		// decode line secret and add to full secret
		decodedSecret, err := hexr.decode([]byte(secret))
		if err != nil {
			return nil, err
		}

		decodedShares[i] = decodedSecret
	}

	result, err := shamir.Combine(decodedShares)
	if err != nil {
		return nil, err
	}

	return result, nil
}
