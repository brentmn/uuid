package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type UUID string

func New() UUID {
	b := make([]byte, 16)
	n, err := rand.Read(b)
	if err != nil {
		panic(fmt.Errorf("uuid: %s", err.Error()))
	}
	if n != 16 {
		panic(fmt.Errorf("uuid: invalid length %d, expecting 16", n))
	}
	b[8] = (b[8] | 0x80) & 0xBF
	b[6] = (b[6] | 0x40) & 0x4F
	return UUID(b)
}

func FromString(s string) (UUID, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("uuid: %s", err.Error())
	}
	if len(b) != 16 {
		return "", fmt.Errorf("uuid: invalid length %d, expecting 32", len(b)*2)
	}
	return UUID(b), nil
}

func (u UUID) String() string {
	return hex.EncodeToString([]byte(u))
}

func (u UUID) MarshalText() ([]byte, error) {
	b := make([]byte, hex.EncodedLen(len(u)))
	hex.Encode(b, []byte(u))
	return b, nil
}

func (u *UUID) UnmarshalText(text []byte) error {
	b := make([]byte, len(text)/2)
	_, err := hex.Decode(b, text)
	if err == nil {
		*u = UUID(b)
	}
	return err
}
