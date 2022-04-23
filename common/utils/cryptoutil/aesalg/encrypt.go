package aesalg

import (
	"c2c/common/utils/cryptoutil"
	"crypto/aes"
	"crypto/cipher"
)

// AES CBC mode Encrypt function
func CBCEncrypt(plaintext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, iv[:blockSize])
	plaintext, err = cryptoutil.Pkcs7Pad(plaintext, blockSize)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}
