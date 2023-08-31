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

	a := abe.NewFAME()                             // Create the scheme instance
	pubKey, secKey, _ := a.GenerateMasterKeys()    // Create a public key and a master secret key
	msp, _ := abe.BooleanToMSP(policy, false)      // The MSP structure defining the policy
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

//MA-ABE
func TestMAABE(t *testing.T) {
	maabe := abe.NewMAABE()

	// create three authorities, each with two attributes
	attribs1 := []string{"auth1:at1", "auth1:at2"}
	attribs2 := []string{"auth2:at1", "auth2:at2"}
	attribs3 := []string{"auth3:at1", "auth3:at2"}
	auth1, err := maabe.NewMAABEAuth("auth1", attribs1)
	if err != nil {
		t.Fatalf("Failed generation authority %s: %v\n", "auth1", err)
	}
	auth2, err := maabe.NewMAABEAuth("auth2", attribs2)
	if err != nil {
		t.Fatalf("Failed generation authority %s: %v\n", "auth2", err)
	}
	auth3, err := maabe.NewMAABEAuth("auth3", attribs3)
	if err != nil {
		t.Fatalf("Failed generation authority %s: %v\n", "auth3", err)
	}

	// create a msp struct out of the boolean formula
	msp, err := abe.BooleanToMSP("((auth1:at1 AND auth2:at1) OR (auth1:at2 AND auth2:at2)) OR (auth3:at1 AND auth3:at2)", false)
	if err != nil {
		t.Fatalf("Failed to generate the policy: %v\n", err)
	}

	// define the set of all public keys we use
	pks := []*abe.MAABEPubKey{auth1.PubKeys(), auth2.PubKeys(), auth3.PubKeys()}

	// choose a message
	msg := "Attack at dawn!"

	// encrypt the message with the decryption policy in msp
	ct, err := maabe.Encrypt(msg, msp, pks)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v\n", err)
	}

	// choose a single user's Global ID
	gid := "gid1"

	// authority 1 issues keys to user
	keys1, err := auth1.GenerateAttribKeys(gid, attribs1)
	if err != nil {
		t.Fatalf("Failed to generate attribute keys: %v\n", err)
	}
	key11, _ := keys1[0], keys1[1]
	// authority 2 issues keys to user
	keys2, err := auth2.GenerateAttribKeys(gid, attribs2)
	if err != nil {
		t.Fatalf("Failed to generate attribute keys: %v\n", err)
	}
	key21, _ := keys2[0], keys2[1]
	// authority 3 issues keys to user
	keys3, err := auth3.GenerateAttribKeys(gid, attribs3)
	if err != nil {
		t.Fatalf("Failed to generate attribute keys: %v\n", err)
	}
	key31, _ := keys3[0], keys3[1]

	// user tries to decrypt with different key combos
	ks1 := []*abe.MAABEKey{key11, key21, key31} // ok

	msg1, err := maabe.Decrypt(ct, ks1)
	if err != nil {
		t.Fatalf("Error decrypting with keyset 1: %v\n", err)
	}
	println(msg1)

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
