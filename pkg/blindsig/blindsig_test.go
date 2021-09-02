package blindsig

import (
	"testing"

	"crypto/rand"
	"crypto/rsa"

	"github.com/cryptoballot/rsablind"
)

const (
	message = "loa"
)

var (
	// the broker's pk-psk
	testKey, _ = rsa.GenerateKey(rand.Reader, keySize)
	// TBD use external key pairs
)

func TestGenSignToken(t *testing.T) {
	b, ub, hashed, err := Gen(message, &testKey.PublicKey)
	if err != nil {
		t.Fatalf("%v", err)
	}
	sig, err := Sign(b, testKey)
	if err != nil {
		t.Fatalf("%v", err)
	}

	// Unblind the signature
	// Note: run at the bTelco
	unblindedSig := rsablind.Unblind(&testKey.PublicKey, sig, ub)

	// Verify the original hashed message against the unblinded signature
	// Note: one can verify without the private signing key
	// Note: the signer has never seen the hashed message and the verifier
	// won't see the original signature.
	ok, err := Verify(hashed, unblindedSig, &testKey.PublicKey)
	if !ok || err != nil {
		t.Fatalf("%v %v", ok, err)
	}
}

// Offline
func BenchmarkGenToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _, _, _ = Gen(message, &testKey.PublicKey)
	}
}

// Offline
func BenchmarkSignToken(b *testing.B) {
	t, _, _, _ := Gen(message, &testKey.PublicKey)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = Sign(t, testKey)
	}
}

// During attachment
func BenchmarkVerifyToken(b *testing.B) {
	t, ub, hashed, _ := Gen(message, &testKey.PublicKey)
	sig, _ := Sign(t, testKey)
	ubsig := rsablind.Unblind(&testKey.PublicKey, sig, ub)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = Verify(hashed, ubsig, &testKey.PublicKey)
	}
}
