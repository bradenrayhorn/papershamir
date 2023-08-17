package papershamir

import "crypto/rand"

func RandomPassphrase() (string, error) {
	random := make([]byte, 10)
	if _, err := rand.Read(random); err != nil {
		return "", err
	}

	return formatIntoBlocks(string(hexr.encode(random))), nil
}
