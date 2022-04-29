package eccalg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

// 参考链接
// https://gofuncchan.gitee.io/books/2-Go-Advance/08-Go-%E5%AF%86%E7%A0%81%E5%AD%A6%E4%BA%94%E9%9D%9E%E5%AF%B9%E7%A7%B0%E5%8A%A0%E5%AF%86%E4%B9%8B%E6%A4%AD%E5%9C%86%E6%9B%B2%E7%BA%BF%E7%AE%97%E6%B3%95/
// https://blog.netlab.360.com/pinkbot/
func GenerateECCKey() ([]byte, []byte, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	pubKey := privateKey.PublicKey
	x, y := pubKey.X, pubKey.Y
	pubKeyBytes := elliptic.Marshal(pubKey.Curve, x, y)

	return privateKey.D.Bytes(), pubKeyBytes, nil
}
