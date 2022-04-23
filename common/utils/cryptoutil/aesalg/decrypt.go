package aesalg

import (
	"c2c/common/utils/cryptoutil"
	"crypto/aes"
	"crypto/cipher"
)

// AES CBC mode Decrypt function
func CBCDecrypt(ciphertext []byte, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	plaintext := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(plaintext, ciphertext)
	plaintext, err = cryptoutil.Pkcs7Unpad(plaintext, blockSize)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
