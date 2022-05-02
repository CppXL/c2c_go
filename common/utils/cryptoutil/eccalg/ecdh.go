package eccalg

import (
	"crypto"
	"fmt"

	"github.com/aead/ecdh"
)

// func implements ECDH algorithm to generaic a shared secret
func ECDHX25519Exchange(privateKey *crypto.PrivateKey, publicKey *crypto.PublicKey) []byte {
	c25519 := ecdh.X25519()
	c25519.Params()
	secret := c25519.ComputeSecret(*privateKey, *publicKey)
	fmt.Println(c25519.Params().BitSize)
	return secret
}
