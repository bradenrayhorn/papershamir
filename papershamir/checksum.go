package papershamir

import (
	"encoding/binary"
	"hash/crc32"
)

type checksumImpl struct{}

var checksum = checksumImpl{}

func (c checksumImpl) create(secret string) []byte {
	checksum := crc32.ChecksumIEEE([]byte(secret))

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, checksum)
	return b
}

func (c checksumImpl) verify(secret string, checksum []byte) bool {
	if len(checksum) != 4 {
		return false
	}
	secretChecksum := crc32.ChecksumIEEE([]byte(secret))

	return secretChecksum == binary.LittleEndian.Uint32([]byte(checksum))
}
