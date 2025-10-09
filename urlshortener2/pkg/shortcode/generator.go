package shortcode

import (
	"crypto/rand"
	"math/big"
)

type Generator interface {
	Generate() (string, error)
}

type Base62Generator struct {
	length int
}

func NewBase62Generator() *Base62Generator {
	return &Base62Generator{length: 8}
}

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func (g *Base62Generator) Generate() (string, error) {
	result := make([]byte, g.length)

	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(62))
		if err != nil {
			return "", err
		}
		result[i]= base62Chars[num.Int64()]
	}
	return string(result), nil
}
