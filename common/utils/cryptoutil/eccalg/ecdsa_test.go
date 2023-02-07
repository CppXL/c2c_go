package eccalg

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"testing"
)

func TestECDSASign(t *testing.T) {
	pubkeyCurve := elliptic.P256()
	privatekey := new(ecdsa.PrivateKey)
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var msg []byte = []byte("hello world")
	res := md5.Sum(msg)
	signhash := res[:]
	signature, err := ECDSASign(privatekey, signhash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r := big.NewInt(0).SetBytes(signature[0:32])
	s := big.NewInt(0).SetBytes(signature[32:64])
	verifystatus := ecdsa.Verify(&privatekey.PublicKey, signhash, r, s)
	fmt.Println(verifystatus)
}
