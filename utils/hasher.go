package utils

import (
	"errors"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	chars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length = uint64(len(chars))
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetUniqueCode() string {
	var hashed string
	for i := 0; i < 10; i++ {
		hashed += string(chars[rand.Intn(len(chars))])
	}
	return hashed
}

func Encode(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(chars[(number % length)])
	}

	return encodedBuilder.String()
}

func Decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		position := strings.IndexRune(chars, symbol)
		if position == -1 {
			return uint64(position), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(position) + uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
