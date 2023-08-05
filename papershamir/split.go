package papershamir

import (
	"fmt"

	"github.com/hashicorp/vault/shamir"
)

const maxLineSize = 30

func Split(secret []byte, parts int, threshold int) ([]string, error) {
	shares, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		return nil, err
	}

	formattedShares := make([]string, parts)
	for j, share := range shares {
		formattedShare := ""

		// convert share to hexr format
		encodedShare := string(hexr.encode(share))

		padding := maxLineSize + (maxLineSize / 2)
		minimalPadding := len(encodedShare) + len(encodedShare)/2
		if padding > minimalPadding {
			padding = minimalPadding
		}

		printLine := func(content string, checksum []byte) string {
			return fmt.Sprintf("%-"+fmt.Sprint(padding)+"s%s\n", content, checksum)
		}

		// split share into lines of length `maxLineSize`
		for i := 0; i < len(encodedShare); i += maxLineSize {
			secretLine := encodedShare[i:]
			if i+maxLineSize <= len(encodedShare) {
				secretLine = encodedShare[i : i+maxLineSize]
			}

			// create checksum for line
			lineChecksum := hexr.encode(checksum.create(secretLine))
			formattedShare += printLine(addWhitespaceToLine(secretLine), lineChecksum)
		}

		// create final checksum for entire share
		shareChecksum := hexr.encode(checksum.create(encodedShare))
		formattedShare += printLine("", shareChecksum)

		formattedShares[j] = formattedShare
	}

	return formattedShares, nil
}

func addWhitespaceToLine(line string) string {
	for i := 2; i < len(line); i += 3 {
		line = line[:i] + " " + line[i:]
	}
	return line
}
