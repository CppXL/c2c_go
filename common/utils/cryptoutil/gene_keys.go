package cryptoutil

import (
	"crypto/rand"
)

// generate aes crypto key return 256bits data
func GenerateAesKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// generate AES CBC mode init vector return 256bits data
func GenerateIv() ([]byte, error) {
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	if err != nil {
		return nil, err
	}
	return iv, nil
}
