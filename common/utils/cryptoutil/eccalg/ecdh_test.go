package eccalg

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/aead/ecdh"
)

func TestECDHExchange(t *testing.T) {
	c25519 := ecdh.X25519()

	privateAlice, publicAlice, err := c25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Alice's private/public key pair: %s\n", err)
	}

	privateBob, publicBob, err := c25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Failed to generate Bob's private/public key pair: %s\n", err)
	}

	if err := c25519.Check(publicBob); err != nil {
		fmt.Printf("Bob's public key is not on the curve: %s\n", err)
	}
	if err := c25519.Check(publicAlice); err != nil {
		fmt.Printf("Alice's public key is not on the curve: %s\n", err)
	}

	s1 := ECDHX25519Exchange(&privateAlice, &publicBob)
	s2 := ECDHX25519Exchange(&privateBob, &publicAlice)
	if !bytes.Equal(s1, s2) {
		fmt.Printf("key exchange failed - secret X coordinates not equal\n")
	}
	t.Log(publicAlice)
	t.Log(publicBob)
	t.Log(privateAlice)
	t.Log(privateBob)
	t.Log("\n" + hex.Dump(s1))

}
