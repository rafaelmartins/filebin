package id

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

const (
	chars    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lenChars = int64(len(chars))
)

func getRandomChar() (uint8, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(lenChars))
	if err != nil {
		return 0, err
	}
	v := n.Int64()
	if v >= lenChars {
		return 0, fmt.Errorf("id: bad random integer: %d", v)
	}
	return chars[v], nil
}

func Generate(length uint8) (string, error) {
	if length == 0 {
		return "", errors.New("id: invalid id length")
	}
	var t strings.Builder
	t.Grow(int(length))
	for i := uint8(0); i < length; i++ {
		r, err := getRandomChar()
		if err != nil {
			return "", err
		}
		if err := t.WriteByte(r); err != nil {
			return "", err
		}
	}
	return t.String(), nil
}
