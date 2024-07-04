package test

import (
	"github.com/fentec-project/gofe/abe"
	"github.com/fentec-project/gofe/data"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

//CP-ABE
func TestCPABE(t *testing.T) {
	msg := "Attack at dawn!"
	policy := "((0 AND 1) OR (2 AND 3)) AND 5"

	gamma := []string{"0", "2", "3", "5"} // owned attributes

	a := abe.NewFAME()                          // Create the scheme instance
	pubKey, secKey, _ := a.GenerateMasterKeys() // Create a public key and a master secret key
	msp, _ := abe.BooleanToMSP(policy, false)
	cipher, _ := a.Encrypt(msg, msp, pubKey)       // Encrypt msg with policy msp under public key pubKey
	keys, _ := a.GenerateAttribKeys(gamma, secKey) // Generate keys for the entity with attributes gamma
	dec, _ := a.Decrypt(cipher, keys, pubKey)      // Decrypt the message
	println(dec)
}

//KP-ABE
func TestKPABE(t *testing.T) {
	l := 10
	a := abe.NewGPSW(l)

	pubKey, secKey, err := a.GenerateMasterKeys()
	if err != nil {
		t.Fatalf("Failed to generate master keys: %v", err)
	}

	msg1 := "Attack at dawn!"

	gamma1 := []int{0, 4, 5}

	cipher1, err := a.Encrypt(msg1, gamma1, pubKey)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	msp, err := abe.BooleanToMSP("(1 OR 4) AND (2 OR (0 AND 5))", true)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v", err)
	}

	abeKey, err := a.GeneratePolicyKey(msp, secKey)
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	msgCheck, err := a.Decrypt(cipher1, abeKey)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	println(msgCheck)

}

//DIPPE
func TestDIPPE(t *testing.T) {
	// create a new DIPPE struct, choosing the security parameter
	d, err := abe.NewDIPPE(3)
	if err != nil {
		t.Fatalf("Failed to generate a new scheme: %v", err)
	}
	vecLen := 5

	// create authorities and their public keys
	auth := make([]*abe.DIPPEAuth, vecLen)
	pubKeys := make([]*abe.DIPPEPubKey, vecLen)
	for i := range auth {
		auth[i], err = d.NewDIPPEAuth(i)
		if err != nil {
			t.Fatalf("Failed to generate a new authority: %v", err)
		}
		pubKeys[i] = &auth[i].Pk
	}

	msg := "some message"

	// choose a policy vector
	policyVec := data.Vector([]*big.Int{big.NewInt(1), big.NewInt(-1),
		big.NewInt(1), big.NewInt(0), big.NewInt(0)})

	// encrypt the message with the chosen policy give by a policy vector,
	cipher, err := d.Encrypt(msg, policyVec, pubKeys)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	// choose a unique user's GID
	userGID := "someGID"
	// choose user's vector, decryption is possible if and only if
	// the users's and policy vector are orthogonal
	userVec := data.Vector([]*big.Int{big.NewInt(0), big.NewInt(1),
		big.NewInt(1), big.NewInt(-3), big.NewInt(6)})

	// authorities generate decryption keys for the user
	userKeys := make([]data.VectorG2, vecLen)
	for i := range auth {
		userKeys[i], err = auth[i].DeriveKeyShare(userVec, pubKeys, userGID)
		if err != nil {
			t.Fatalf("Failed to generate a user key: %v", err)
		}
	}

	// user decrypts using collected keys
	dec, err := d.Decrypt(cipher, userKeys, userVec, userGID)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}
	assert.Equal(t, msg, dec)
}
