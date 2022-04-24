package aesalg

import (
	"c2c/common/utils/cryptoutil"
	"encoding/hex"
	"testing"
)

func TestEncrypt(t *testing.T) {

	enckey, _ := cryptoutil.GenerateAesKey()
	plaintext := []byte("你是谁，我是你爹")
	iv, _ := cryptoutil.GenerateIv()
	ciphertext, err := CBCEncrypt(plaintext, enckey, iv)
	if err != nil {
		t.Error(err)
	}
	t.Log("\n密文\n" + hex.Dump(ciphertext))
	t.Log("开始解密")
	plaintext, err = CBCDecrypt(ciphertext, enckey, iv)
	if err != nil {
		t.Error(err)
	}
	t.Log("明文:" + hex.Dump(plaintext))
	t.Log("\n" + string(plaintext))
}
