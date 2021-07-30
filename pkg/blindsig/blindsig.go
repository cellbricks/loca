package blindsig

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"

	"github.com/cryptoballot/fdh"
	"github.com/cryptoballot/rsablind"
)

const (
	hashSize = 1536
	keySize  = 2048
)

// TBD external key
var key, _ = rsa.GenerateKey(rand.Reader, keySize)

func GenToken(m string) ([]byte, []byte, error) {
	message := []byte(m)

	// We do a SHA256 full-domain-hash expanded to 1536 bits (3/4 the key size)
	hashed := fdh.Sum(crypto.SHA256, hashSize, message)

	// Blind the hashed message
	blinded, unblinder, err := rsablind.Blind(&key.PublicKey, hashed)
	if err != nil {
		panic(err)
	}

	// Blind sign the blinded message
	sig, err := rsablind.BlindSign(key, blinded)
	if err != nil {
		return nil, nil, err
	}
	return sig, unblinder, nil
}
