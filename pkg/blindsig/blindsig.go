package blindsig

import (
	"crypto"
	"crypto/rsa"
	_ "crypto/sha256"

	"github.com/cryptoballot/fdh"
	"github.com/cryptoballot/rsablind"
)

const (
	hashSize = 1536
	keySize  = 2048
)


// Gen runs at the bTelco to generate a blinded message
// using the broker's public key.
func Gen(message string, key *rsa.PublicKey) ([]byte, []byte, []byte, error) {
	mBytes := []byte(message)
	hashed := fdh.Sum(crypto.SHA256, hashSize, mBytes)

	// Blind the hashed message; unblinder -> random padding
	blinded, unblinder, err := rsablind.Blind(key, hashed)
	if err != nil {
		panic(err)
	}

	return blinded, unblinder, hashed, nil
}

// Sign runs at the broker to sign the bTelco generated
// tokens; it returns the (blinded) signature.
func Sign(token []byte, key *rsa.PrivateKey) ([]byte, error) {
	sig, err := rsablind.BlindSign(key, token)
	if err != nil {
		return nil, err

	}
	return sig, nil
}

// Verify runs at the broker to validate the unblinded sig.
func Verify(message []byte, sig []byte, key *rsa.PublicKey) (bool, error) {
	err := rsablind.VerifyBlindSignature(key, message, sig)
	if err != nil {
		return false, err

	}
	return true, nil
}
