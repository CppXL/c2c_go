package eccalg

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

// https://socketloop.com/tutorials/golang-example-for-ecdsa-elliptic-curve-digital-signature-algorithm-functions
// func ECDSA
func ECDSASign(privatekey *ecdsa.PrivateKey, signhash []byte) ([]byte, error) {
	//
	var pubkey ecdsa.PublicKey
	pubkey = privatekey.PublicKey
	fmt.Println("Private Key :")
	fmt.Printf("%x \n", privatekey)

	fmt.Println("Public Key :")
	fmt.Printf("%x \n", pubkey)
	r := big.NewInt(0)
	s := big.NewInt(0)
	var err error
	r, s, err = ecdsa.Sign(rand.Reader, privatekey, signhash)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)

	return signature, nil
}
