package utils

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"golang.org/x/exp/slices"
)

func ReadByte(r io.Reader) (byte, error) {
	b := make([]byte, 1)
	_, err := r.Read(b)

	return b[0], err
}

// / Read8BytesNE reads 8 bytes in Network Endian from a reader.
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

func IpToUint32(ip net.IP) uint32 {
	return binary.BigEndian.Uint32(ip)
}

func Uint32ToIp(n uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, n)
	return ip
}

func InsertSorted[T any](sortedTs []T, t T, compare func(x, y T) int) ([]T, int) {
	i, _ := slices.BinarySearchFunc(sortedTs, t, compare)
	var dummy T
	sortedTs = append(sortedTs, dummy)
	if i < len(sortedTs) {
		copy(sortedTs[i+1:], sortedTs[i:])
		sortedTs[i] = t
	} else {
		sortedTs[i-1] = t
	}
	return sortedTs, i
}

// XorSlice returns xor of two soretd slices.
func XorSlice[T any](s1, s2 []T, compare func(T, T) int) []T {
	out := []T{}
	j := 0

	for _, x := range s1 {
		for jj, y := range s2 {
			res := compare(x, y)
			if jj == j {
				if res < 0 {
					out = append(out, x)
					break
				} else if res == 0 {
					continue
				} else {
					out = append(out, y)
					break
				}
			} else {
				if res < 0 {
					j = jj - 1
					break
				} else if res == 0 {
					continue
				} else {
					// never happens
				}
			}
		}
	}

	return out
}
