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

func RemoveIndex[T any](s []T, i int) []T {
	s = append(s[:i], s[i+1:]...)
	return s
}

func RemoveDuplicates[T any](s []T, compare func(x, y T) int) []T {
	out := make([]T, len(s))
	copy(out, s)

	for i := len(out) - 1; i > 0; i-- {
		x, y := out[i], out[i-1]

		if compare(x, y) == 0 {
			out = RemoveIndex(out, i)
		}
	}

	return out
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

	for {
		for i != 0 && i < len(s1) && compare(s1[i-1], s1[i]) == 0 {
			i++
		}
		for j != 0 && j < len(s2) && compare(s2[j-1], s2[j]) == 0 {
			j++
		}

		if i >= len(s1) {
			log.Printf("j[%v] is greater than len", j)
			i--
			break
		}

		if j >= len(s2) {
			log.Printf("j[%v] is greater than len", j)
			i--
			break
		}

		x, y := s1[i], s2[j]
		res := compare(x, y)

		if res < 0 {
			log.Printf("x %v[%v] < y %v[%v]", x, i, y, j)
			out = append(out, x)
			i += 1
			continue
		} else if 0 < res {
			log.Printf("x %v[%v] > y %v[%v]", x, i, y, j)
			out = append(out, y)
			j += 1
			continue
		}

		log.Printf("x %v[%v] == y %v[%v]", x, i, y, j)
		i++
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
