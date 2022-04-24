package eccalg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

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

func Encrypt() {

}

func Decrypt() {

}
