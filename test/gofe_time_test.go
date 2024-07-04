package test

import (
	"fmt"
	"github.com/fentec-project/gofe/abe"
	"github.com/fentec-project/gofe/data"
	"math/big"
	"testing"
	"time"
)

//CP-ABE
//--------Encrypt time: 23.1232ms
//--------GenerateAttribKeys time: 3.521ms
//--------Decrypt time: 5.3726ms
func TestCPABEUserKey(t *testing.T) {
	msg := "Attack at dawn!"
	policy := "((0 AND 1) OR (2 AND 3)) OR 5"

	gamma := []string{"5"} // owned attributes

	a := abe.NewFAME()
	pubKey, secKey, _ := a.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(policy, false)

	start := time.Now()
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------Encrypt time: %s\n", elapsed)

	start = time.Now()
	keys, _ := a.GenerateAttribKeys(gamma, secKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys time: %s\n", elapsed)

	start = time.Now()
	dec, _ := a.Decrypt(cipher, keys, pubKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
	println(dec)
}

//--------Encrypt time: 23.4988ms
//--------GenerateAttribKeys time: 4.2279ms
//--------Decrypt time: 5.5477ms
func TestCPABEUserKey2(t *testing.T) {
	msg := "Attack at dawn!"
	policy := "((0 AND 1) OR (2 AND 3)) OR 5"

	gamma := []string{"0", "1"} // owned attributes

	a := abe.NewFAME()
	pubKey, secKey, _ := a.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(policy, false)

	start := time.Now()
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------Encrypt time: %s\n", elapsed)

	start = time.Now()
	keys, _ := a.GenerateAttribKeys(gamma, secKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys time: %s\n", elapsed)

	start = time.Now()
	dec, _ := a.Decrypt(cipher, keys, pubKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
	println(dec)
}

//--------Encrypt time: 24.73ms
//--------GenerateAttribKeys time: 5.9582ms
//--------Decrypt time: 5.8354ms
func TestCPABEUserKey3(t *testing.T) {
	msg := "Attack at dawn!"
	policy := "((0 AND 1) OR (2 AND 3)) OR 5"

	gamma := []string{"0", "1", "2"} // owned attributes

	a := abe.NewFAME()
	pubKey, secKey, _ := a.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(policy, false)

	start := time.Now()
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------Encrypt time: %s\n", elapsed)

	start = time.Now()
	keys, _ := a.GenerateAttribKeys(gamma, secKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys time: %s\n", elapsed)

	start = time.Now()
	dec, _ := a.Decrypt(cipher, keys, pubKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
	println(dec)
}

//--------Encrypt time: 24.5175ms
//--------GenerateAttribKeys time: 7.5087ms
//--------Decrypt time: 6.0958ms
func TestCPABEUserKey4(t *testing.T) {
	msg := "Attack at dawn!"
	policy := "((0 AND 1) OR (2 AND 3)) OR 5"

	gamma := []string{"0", "1", "2", "3"} // owned attributes

	a := abe.NewFAME()
	pubKey, secKey, _ := a.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(policy, false)

	start := time.Now()
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------Encrypt time: %s\n", elapsed)

	start = time.Now()
	keys, _ := a.GenerateAttribKeys(gamma, secKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys time: %s\n", elapsed)

	start = time.Now()
	dec, _ := a.Decrypt(cipher, keys, pubKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
	println(dec)
}

//--------Encrypt time: 23.166ms
//--------GenerateAttribKeys time: 8.0594ms
//--------Decrypt time: 0s
func TestCPABEUserKey5(t *testing.T) {
	msg := "Attack at dawn!"
	policy := "((0 AND 1) OR (2 AND 3)) OR 5"

	gamma := []string{"1", "2", "3", "4", "5"} // owned attributes

	a := abe.NewFAME()
	pubKey, secKey, _ := a.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(policy, false)

	start := time.Now()
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------Encrypt time: %s\n", elapsed)

	start = time.Now()
	keys, _ := a.GenerateAttribKeys(gamma, secKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys time: %s\n", elapsed)

	start = time.Now()
	dec, _ := a.Decrypt(cipher, keys, pubKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
	println(dec)
}

//--------Encrypt time: 26.8878ms
//--------GenerateAttribKeys time: 9.7068ms
//--------Decrypt time: 5.0429ms
func TestCPABEUserKey6(t *testing.T) {
	msg := "Attack at dawn!"
	policy := "((0 AND 1) OR (2 AND 3)) OR 5"

	gamma := []string{"0", "1", "2", "3", "4", "5"} // owned attributes

	a := abe.NewFAME()
	pubKey, secKey, _ := a.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(policy, false)

	start := time.Now()
	cipher, _ := a.Encrypt(msg, msp, pubKey)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------Encrypt time: %s\n", elapsed)

	start = time.Now()
	keys, _ := a.GenerateAttribKeys(gamma, secKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys time: %s\n", elapsed)

	start = time.Now()
	dec, _ := a.Decrypt(cipher, keys, pubKey)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
	println(dec)
}

//MA-ABE
//--------GenerateAttribKeys 1: 1.7761ms
//--------GenerateAttribKeys 2: 3.4919ms
//--------GenerateAttribKeys 3: 5.1312ms
//--------GenerateAttribKeys 4: 7.1995ms
//--------GenerateAttribKeys 5: 9.3941ms
//--------GenerateAttribKeys 6: 11.5877ms
//--------Decrypt time 1 : 2.148ms
//--------Decrypt time 2 : 3.1401ms
//--------Decrypt time 3 : 4.5377ms
//--------Decrypt time 4 : 7.6194ms
//--------Decrypt time 5 : 8.8556ms
//--------Decrypt time 6 : 11.6404ms
func TestMAABE(t *testing.T) {
	maabe := abe.NewMAABE()

	// create three authorities, each with two attributes
	attribs1 := []string{"auth1:at1", "auth1:at2"}
	attribs2 := []string{"auth2:at1", "auth2:at2"}
	attribs3 := []string{"auth3:at1", "auth3:at2"}
	attribs4 := []string{"auth4:at1", "auth4:at2"}
	attribs5 := []string{"auth5:at1", "auth5:at2"}
	attribs6 := []string{"auth6:at1", "auth6:at2"}

	start := time.Now()
	auth1, err := maabe.NewMAABEAuth("auth1", attribs1)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys 1: %s\n", elapsed)
	if err != nil {
		t.Fatalf("Failed generation authority %s: %v\n", "auth1", err)
	}
	auth2, err := maabe.NewMAABEAuth("auth2", attribs2)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys 2: %s\n", elapsed)
	if err != nil {
		t.Fatalf("Failed generation authority %s: %v\n", "auth2", err)
	}
	auth3, err := maabe.NewMAABEAuth("auth3", attribs3)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys 3: %s\n", elapsed)
	if err != nil {
		t.Fatalf("Failed generation authority %s: %v\n", "auth3", err)
	}
	_, err = maabe.NewMAABEAuth("auth4", attribs4)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys 4: %s\n", elapsed)
	_, err = maabe.NewMAABEAuth("auth5", attribs5)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys 5: %s\n", elapsed)
	_, err = maabe.NewMAABEAuth("auth6", attribs6)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys 6: %s\n", elapsed)

	// create a msp struct out of the boolean formula
	msp, err := abe.BooleanToMSP("((auth1:at1 OR auth1:at2 ) OR (auth2:at1AND auth2:at2)) OR (auth3:at1 AND auth3:at2)", false)
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

	start = time.Now()
	// authority 1 issues keys to user
	keys1, err := auth1.GenerateAttribKeys(gid, attribs1)
	if err != nil {
		t.Fatalf("Failed to generate attribute keys: %v\n", err)
	}
	key11, key12 := keys1[0], keys1[1]
	// authority 2 issues keys to user
	keys2, err := auth2.GenerateAttribKeys(gid, attribs2)
	if err != nil {
		t.Fatalf("Failed to generate attribute keys: %v\n", err)
	}
	key21, key22 := keys2[0], keys2[1]
	//authority 3 issues keys to user
	keys3, err := auth3.GenerateAttribKeys(gid, attribs3)
	if err != nil {
		t.Fatalf("Failed to generate attribute keys: %v\n", err)
	}
	key31, key32 := keys3[0], keys3[1]

	// user tries to decrypt with different key combos
	ks1 := []*abe.MAABEKey{key11, key12, key21, key22} // ok
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------GenerateAttribKeys: %s\n", elapsed)

	start = time.Now()
	msg1, err := maabe.Decrypt(ct, ks1)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time : %s\n", elapsed)

	if err != nil {
		t.Fatalf("Error decrypting with keyset 1: %v\n", err)
	}
	println(msg1)

	start = time.Now()
	ks_1 := []*abe.MAABEKey{key11}
	_, _ = maabe.Decrypt(ct, ks_1)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time 1 : %s\n", elapsed)

	start = time.Now()
	ks_2 := []*abe.MAABEKey{key11, key12}
	_, _ = maabe.Decrypt(ct, ks_2)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time 2 : %s\n", elapsed)

	start = time.Now()
	ks_3 := []*abe.MAABEKey{key11, key12, key21}
	_, _ = maabe.Decrypt(ct, ks_3)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time 3 : %s\n", elapsed)

	start = time.Now()
	ks_4 := []*abe.MAABEKey{key11, key12, key21, key22}
	_, _ = maabe.Decrypt(ct, ks_4)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time 4 : %s\n", elapsed)

	start = time.Now()
	ks_5 := []*abe.MAABEKey{key11, key12, key21, key22, key31}
	_, _ = maabe.Decrypt(ct, ks_5)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time 5 : %s\n", elapsed)

	start = time.Now()
	ks_6 := []*abe.MAABEKey{key11, key12, key21, key22, key31, key32}
	_, _ = maabe.Decrypt(ct, ks_6)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time 6 : %s\n", elapsed)
}

//DIPPE
//--------Generate UserKey [0]: 11.5275ms
//--------Generate UserKey [1]: 25.2239ms
//--------Generate UserKey [2]: 37.5642ms
//--------Generate UserKey [3]: 50.0152ms
//--------Generate UserKey [4]: 62.1566ms
//--------Generate UserKey [5]: 74.4976ms
//--------Generate UserKey [6]: 86.014ms
//--------Decrypt time: 8.2942ms
func TestDIPPEUserKey(t *testing.T) {
	d, err := abe.NewDIPPE(3)
	if err != nil {
		t.Fatalf("Failed to generate a new scheme: %v", err)
	}
	vecLen := 7

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

	policyVec := data.Vector([]*big.Int{big.NewInt(1), big.NewInt(-1),
		big.NewInt(1), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)})

	cipher, err := d.Encrypt(msg, policyVec, pubKeys)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	userGID := "someGID"
	userVec := data.Vector([]*big.Int{big.NewInt(0), big.NewInt(1),
		big.NewInt(1), big.NewInt(-3), big.NewInt(6), big.NewInt(6), big.NewInt(6)})

	start := time.Now()
	userKeys := make([]data.VectorG2, vecLen)
	for i := range auth {
		userKeys[i], err = auth[i].DeriveKeyShare(userVec, pubKeys, userGID)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("--------Generate UserKey [%v]: %s\n", i, elapsed)
		if err != nil {
			t.Fatalf("Failed to generate a user key: %v", err)
		}
		//start2 := time.Now()
		//_, _ = d.Decrypt(cipher, userKeys, userVec, userGID)
		//end2 := time.Now()
		//elapsed = end2.Sub(start2)
		//fmt.Printf("--------Decrypt time[%v]: %s\n", i, elapsed)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------GeneratePolicyKey time: %s\n", elapsed)

	start = time.Now()
	_, _ = d.Decrypt(cipher, userKeys, userVec, userGID)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
}

func TestDIPPEUserKey2(t *testing.T) {
	d, err := abe.NewDIPPE(3)
	if err != nil {
		t.Fatalf("Failed to generate a new scheme: %v", err)
	}
	vecLen := 3

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

	policyVec := data.Vector([]*big.Int{big.NewInt(1), big.NewInt(-1),
		big.NewInt(1)})

	cipher, err := d.Encrypt(msg, policyVec, pubKeys)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	userGID := "someGID"
	userVec := data.Vector([]*big.Int{big.NewInt(0), big.NewInt(1),
		big.NewInt(1)})

	start := time.Now()
	userKeys := make([]data.VectorG2, vecLen)
	for i := range auth {
		userKeys[i], err = auth[i].DeriveKeyShare(userVec, pubKeys, userGID)
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("--------Generate UserKey [%v]: %s\n", i, elapsed)
		if err != nil {
			t.Fatalf("Failed to generate a user key: %v", err)
		}
		//start2 := time.Now()
		//_, _ = d.Decrypt(cipher, userKeys, userVec, userGID)
		//end2 := time.Now()
		//elapsed = end2.Sub(start2)
		//fmt.Printf("--------Decrypt time[%v]: %s\n", i, elapsed)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------GeneratePolicyKey time: %s\n", elapsed)

	start = time.Now()
	_, _ = d.Decrypt(cipher, userKeys, userVec, userGID)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decrypt time: %s\n", elapsed)
}
