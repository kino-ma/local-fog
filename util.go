package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

func ReadByte(r io.Reader) (byte, error) {
	b := make([]byte, 1)
	_, err := r.Read(b)

	return b[0], err
}

/// Read8BytesNE reads 8 bytes in Network Endian from a reader.
func Read8BytesNE(r io.Reader) (uint64, error) {
	b := make([]byte, 8)
	n, err := r.Read(b)

	if n < 8 {
		err = fmt.Errorf("insufficient bytes read: %v", n)
		return 0, err
	}

	result := binary.BigEndian.Uint64(b)

	return result, err
}
