package papershamir

import (
	"github.com/hashicorp/vault/shamir"
)

const maxLineSize = 40

func Split(secret []byte, parts int, threshold int, key string) ([]string, error) {
	shares, err := splitBytes(secret, parts, threshold, key)
	if err != nil {
		return nil, err
	}

	formattedShares := make([]string, parts)
	for j, share := range shares {
		formattedShare := ""

		// convert share to hexr format
		encodedShare := string(hexr.encode(share))

		// split share into lines of length `maxLineSize`
		for i := 0; i < len(encodedShare); i += maxLineSize {
			secretLine := encodedShare[i:]
			if i+maxLineSize <= len(encodedShare) {
				secretLine = encodedShare[i : i+maxLineSize]
			}

			// create checksum for line
			lineChecksum := string(hexr.encode(checksum.create(secretLine)))
			formattedShare += addWhitespaceToLine(secretLine + lineChecksum)
		}

		// create final checksum for entire share
		shareChecksum := hexr.encode(checksum.create(encodedShare))
		formattedShare += addWhitespaceToLine(string(shareChecksum))

		formattedShares[j] = formattedShare
	}

	return formattedShares, nil
}

func splitBytes(secret []byte, parts int, threshold int, key string) ([][]byte, error) {
	if key != "" {
		encryptedSecret, err := encrypt.encrypt(key, secret)
		if err != nil {
			return nil, err
		}
		secret = encryptedSecret
	}

	shares, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		return nil, err
	}

	return shares, err
}

func addWhitespaceToLine(line string) string {
	return formatIntoBlocks(line) + "\n"
}
