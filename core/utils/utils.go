package utils

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
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
	i, j := 0, 0
	var x T

	for i, x = range s1 {
		if j >= len(s2) {
			i--
			break
		}

		y := s2[j]
		res := compare(x, y)

		if res < 0 {
			out = append(out, x)
			continue
		} else if 0 < res {
			out = append(out, y)
			j += 1
			continue
		}

		for j+1 < len(s2) && compare(s2[j], s2[j+1]) == 0 {
			j++
		}

		j++
	}

	log.Print(i, j, out)

	for _, x := range s1[i+1:] {
		out = append(out, x)
	}

	for _, y := range s2[j:] {
		out = append(out, y)
	}

	return out
}
