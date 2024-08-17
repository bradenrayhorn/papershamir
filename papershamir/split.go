package papershamir

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"

	"github.com/hashicorp/vault/shamir"
	"github.com/klauspost/compress/zstd"
)

const maxLineSize = 40

const prefixCompressed = byte('1')
const prefixUncompressed = byte('2')

func SplitHexr(secret []byte, parts int, threshold int, key string) ([]string, error) {
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

func SplitQR(secret []byte, parts int, threshold int, key string) ([]string, error) {
	shares, err := splitBytes(secret, parts, threshold, key)
	if err != nil {
		return nil, err
	}

	formattedShares := make([]string, parts)
	for j, share := range shares {
		formattedShares[j] = string(base64.RawStdEncoding.EncodeToString(share))
	}

	return formattedShares, nil
}

func splitBytes(secret []byte, parts int, threshold int, key string) ([][]byte, error) {
	// maybe compress
	var b bytes.Buffer
	encoder, err := zstd.NewWriter(&b, zstd.WithEncoderLevel(zstd.SpeedBestCompression))
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(encoder, bytes.NewReader(secret))
	if err != nil {
		_ = encoder.Close()
		return nil, err
	}
	if err := encoder.Close(); err != nil {
		return nil, err
	}

	compressedSecret := b.Bytes()
	if len(secret) == 0 {
		return nil, errors.New("cannot split an empty secret")
	} else if len(compressedSecret) < len(secret) {
		secret = append([]byte{prefixCompressed}, compressedSecret...)
	} else {
		secret = append([]byte{prefixUncompressed}, secret...)
	}

	// encrypt
	if key != "" {
		encryptedSecret, err := encrypt.encrypt(key, secret)
		if err != nil {
			return nil, err
		}
		secret = encryptedSecret
	}

	// split
	shares, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		return nil, err
	}

	return shares, err
}

func addWhitespaceToLine(line string) string {
	return formatIntoBlocks(line) + "\n"
}
