package papershamir

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/vault/shamir"
	"github.com/klauspost/compress/zstd"
)

func CombineHexr(shares [][]byte, key string) ([]byte, error) {
	decodedShares := make([][]byte, len(shares))

	for i, share := range shares {
		lines := strings.Split(strings.TrimSpace(string(share)), "\n")

		if len(lines) < 2 {
			return nil, fmt.Errorf("share requires at least two lines. at share %d", i+1)
		}

		secret := ""
		for j, line := range lines[:len(lines)-1] {
			stripped := strings.ReplaceAll(line, " ", "")
			if len(stripped) < 10 {
				return nil, fmt.Errorf("invalid format. at share %d line %d", i+1, j+1)
			}

			foundChecksum := stripped[len(stripped)-8:]
			foundSecret := stripped[:len(stripped)-8]
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
		shareChecksum := strings.TrimSpace(strings.ReplaceAll(lines[len(lines)-1], " ", ""))
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

	result, err := combineBytes(decodedShares, key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CombineQR(shares [][]byte, key string) ([]byte, error) {
	decodedShares := make([][]byte, len(shares))

	for i, share := range shares {
		share := strings.TrimSpace(string(share))
		byteShare, err := base64.RawStdEncoding.DecodeString(share)
		if err != nil {
			return nil, err
		}

		decodedShares[i] = byteShare
	}

	result, err := combineBytes(decodedShares, key)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func combineBytes(shares [][]byte, key string) ([]byte, error) {
	// combine
	result, err := shamir.Combine(shares)
	if err != nil {
		return nil, err
	}

	// decrypt
	if key != "" {
		result, err = encrypt.decrypt(key, result)
		if err != nil {
			return nil, err
		}
	}

	// decompress
	if len(result) == 0 {
		return nil, errors.New("secret is empty")
	}

	prefix := result[0]
	result = result[1:]
	if prefix == prefixCompressed {
		var b bytes.Buffer
		decoder, err := zstd.NewReader(bytes.NewReader(result))
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(&b, decoder)
		decoder.Close()
		if err != nil {
			return nil, err
		}
		result = b.Bytes()
	}

	return result, nil
}
