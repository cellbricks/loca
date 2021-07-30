package blindsig

import (
	"github.com/cryptoballot/rsablind"
	"testing"
)

func TestGenToken(t *testing.T) {
	padding := "loa"
	want := "loa"
	token, ub, err := GenToken(padding)

	if err != nil {
		t.Fatalf(`Token("loa") = %q, %v, want match for %#q, nil`, token, err, want)
	}

	// Unblind the signature
	unblindedSig := rsablind.Unblind(&key.PublicKey, sig, unblinder)

	// Verify the original hashed message against the unblinded signature
	if err := rsablind.VerifyBlindSignature(&key.PublicKey, hashed, unblindedSig); err != nil {
		t.Fatalf(`Token("loa") = %q, %v, want match for %#q, nil`, token, err, want)
	} else {
		fmt.Println("ALL IS WELL")
	}
}
