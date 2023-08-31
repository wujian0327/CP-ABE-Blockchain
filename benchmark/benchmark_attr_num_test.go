package benchmark

import (
	"CP-ABE-Blockchain/core"
	"CP-ABE-Blockchain/tools"
	"encoding/json"
	"fmt"
	"github.com/fentec-project/gofe/abe"
	"github.com/fentec-project/gofe/data"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

//CPABE Attr5 test
//8         136554812 ns/op
//dtObjectBytesSize: 10436
func BenchmarkCPABEAttrEncrypt_5(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((A AND B) OR (C AND D) OR E)"
	inst := abe.NewFAME()
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", accessStructure)
		CPABEShareKeyToBlockchain(inst, key)
	}
}

//CPABE Attr10 test
//6         166798367 ns/op
//dtObjectBytesSize: 16295
func BenchmarkCPABEAttrEncrypt_10(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J)"
	inst := abe.NewFAME()
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", accessStructure)
		CPABEShareKeyToBlockchain(inst, key)
	}
}

//CPABE Attr15 test
// 5         218277420 ns/op
//dtObjectBytesSize: 22194
func BenchmarkCPABEAttrEncrypt_15(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O)"
	inst := abe.NewFAME()
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", accessStructure)
		CPABEShareKeyToBlockchain(inst, key)
	}
}

//CPABE Attr20 test
//4         289363150 ns/op
//dtObjectBytesSize: 28141
func BenchmarkCPABEAttrEncrypt_20(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O" +
		" OR (P AND Q) OR (R AND S) OR T)"
	inst := abe.NewFAME()
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", accessStructure)
		CPABEShareKeyToBlockchain(inst, key)
	}
}

//CPABE Attr25 test
//3         396678333 ns/op
//dtObjectBytesSize: 34168
func BenchmarkCPABEAttrEncrypt_25(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O" +
		" OR (P AND Q) OR (R AND S) OR T OR (U AND V) OR (W AND X) OR Y)"
	inst := abe.NewFAME()
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", accessStructure)
		CPABEShareKeyToBlockchain(inst, key)
	}
}

//CPABE Attr30 test
// 3         468285567 ns/op
//dtObjectBytesSize: 40236
func BenchmarkCPABEAttrEncrypt_30(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O" +
		" OR (P AND Q) OR (R AND S) OR T OR (U AND V) OR (W AND X) OR Y" +
		" OR (Z AND AA) OR (AB AND AC) OR AD )"
	inst := abe.NewFAME()
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", accessStructure)
		CPABEShareKeyToBlockchain(inst, key)
	}
}

//MAABE Attr5
//7         177639614 ns/op
//dtObjectBytesSize: 17878
func BenchmarkMAABEAttr_5(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth2:at1) OR (auth1:at2 AND auth2:at2)) OR (auth3:at1)"
	maabe := abe.NewMAABE()

	for i := 0; i < b.N; i++ {

		core.CreateDTObject("td_sample.json", accessStructure)
		MAABEShareKeyToBlockchain(maabe, key)
	}

}

//MAABE Attr10
//6         214656917 ns/op
//dtObjectBytesSize: 33519
func BenchmarkMAABEAttr_10(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5))"
	maabe := abe.NewMAABE()

	for i := 0; i < b.N; i++ {

		core.CreateDTObject("td_sample.json", accessStructure)
		MAABEShareKeyToBlockchain(maabe, key)
	}

}

//MAABE Attr15
//5         231950320 ns/op
//dtObjectBytesSize: 49208
func BenchmarkMAABEAttr_15(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5))"
	maabe := abe.NewMAABE()

	for i := 0; i < b.N; i++ {

		core.CreateDTObject("td_sample.json", accessStructure)
		MAABEShareKeyToBlockchain(maabe, key)
	}

}

//MAABE Attr20
//5         257341200 ns/op
//dtObjectBytesSize: 64902
func BenchmarkMAABEAttr_20(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5)" +
		" OR (auth1:at6 AND auth1:at7) OR (auth1:at8 AND auth1:at9) OR (auth1:at10))"
	maabe := abe.NewMAABE()

	for i := 0; i < b.N; i++ {

		core.CreateDTObject("td_sample.json", accessStructure)
		MAABEShareKeyToBlockchain(maabe, key)
	}

}

//MAABE Attr25
//5         278099380 ns/op
//dtObjectBytesSize: 80704
func BenchmarkMAABEAttr_25(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5)" +
		" OR (auth2:at6 AND auth2:at7) OR (auth2:at8 AND auth2:at9) OR (auth2:at10)" +
		" OR (auth1:at6 AND auth1:at7) OR (auth1:at8 AND auth1:at9) OR (auth1:at10))"
	maabe := abe.NewMAABE()

	for i := 0; i < b.N; i++ {

		core.CreateDTObject("td_sample.json", accessStructure)
		MAABEShareKeyToBlockchain(maabe, key)
	}

}

//MAABE Attr30
//4         308813025 ns/op
//dtObjectBytesSize: 96498
func BenchmarkMAABEAttr_30(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5)" +
		" OR (auth3:at6 AND auth3:at7) OR (auth3:at8 AND auth3:at9) OR (auth3:at10)" +
		" OR (auth2:at6 AND auth2:at7) OR (auth2:at8 AND auth2:at9) OR (auth2:at10)" +
		" OR (auth1:at6 AND auth1:at7) OR (auth1:at8 AND auth1:at9) OR (auth1:at10))"
	maabe := abe.NewMAABE()

	for i := 0; i < b.N; i++ {

		core.CreateDTObject("td_sample.json", accessStructure)
		MAABEShareKeyToBlockchain(maabe, key)
	}

}

//DIPPE Attr30
//4         208154271 ns/op
//dtObjectBytesSize: 11381
func BenchmarkDIPPEAttr_5(b *testing.B) {
	key := tools.RandKey(32)
	// choose a policy vector
	accessStructure := "0,-1,1,0,0"

	dippe, _ := abe.NewDIPPE(3)

	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", accessStructure)
		DIPPEShareKeyToBlockchain(dippe, key)
	}
}

func CPABEShareKeyToBlockchain(pkInst *abe.FAME, key []byte) {
	PK, _, err := pkInst.GenerateMasterKeys()
	//get access structure from blockchain
	dtObject := core.ReadTDToDTObject("td_sample.json")
	dtObjectByte, err := tools.QueryChaincode(core.ChaincodeId, "GetDTObject", dtObject.Id)
	err = json.Unmarshal(dtObjectByte, &dtObject)
	if err != nil {
		panic(err)
	}
	accessStruct := dtObject.AccessStruct
	msp, err := abe.BooleanToMSP(accessStruct, false)
	if err != nil {
		panic(err)
	}
	cipher, err := pkInst.Encrypt(string(key), msp, PK)

	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	dtObject.CT = CTBytes
	dtObject.IPFSAddress = "IPFS_ADDRESS_123456"
	dtObjectBytes, _ := json.Marshal(dtObject)
	//call update smart contract
	result, err := tools.ExecuteChaincode(core.ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	fmt.Printf("dtObjectBytesSize: %v \n", len(dtObjectBytes))
}

func MAABEShareKeyToBlockchain(maabe *abe.MAABE, key []byte) {
	// create three authorities, each with two attributes
	attribs1 := []string{"auth1:at1", "auth1:at2", "auth1:at3", "auth1:at4", "auth1:at5",
		"auth1:at6", "auth1:at7", "auth1:at8", "auth1:at9", "auth1:at10"}
	attribs2 := []string{"auth2:at1", "auth2:at2", "auth2:at3", "auth2:at4", "auth2:at5",
		"auth2:at6", "auth2:at7", "auth2:at8", "auth2:at9", "auth2:at10"}
	attribs3 := []string{"auth3:at1", "auth3:at2", "auth3:at3", "auth3:at4", "auth3:at5",
		"auth3:at6", "auth3:at7", "auth3:at8", "auth3:at9", "auth3:at10"}
	auth1, _ := maabe.NewMAABEAuth("auth1", attribs1)

	auth2, _ := maabe.NewMAABEAuth("auth2", attribs2)

	auth3, _ := maabe.NewMAABEAuth("auth3", attribs3)

	//get access structure from blockchain
	dtObject := core.ReadTDToDTObject("td_sample.json")
	dtObjectByte, _ := tools.QueryChaincode(core.ChaincodeId, "GetDTObject", dtObject.Id)
	_ = json.Unmarshal(dtObjectByte, &dtObject)

	accessStructure := dtObject.AccessStruct
	// create a msp struct out of the boolean formula
	msp, err := abe.BooleanToMSP(accessStructure, false)
	if err != nil {
		panic(err)
	}

	// define the set of all public keys we use
	pks := []*abe.MAABEPubKey{auth1.PubKeys(), auth2.PubKeys(), auth3.PubKeys()}

	// encrypt the message with the decryption policy in msp
	ct, _ := maabe.Encrypt(string(key), msp, pks)

	//save CP and ipfsAddress to blockchain

	CTBytes, _ := json.Marshal(ct)
	dtObject.CT = CTBytes
	dtObject.IPFSAddress = "IPFS_ADDRESS_123456"
	dtObjectBytes, _ := json.Marshal(dtObject)
	//call update smart contract
	result, err := tools.ExecuteChaincode(core.ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	fmt.Printf("dtObjectBytesSize: %v \n", len(dtObjectBytes))
}

func DIPPEShareKeyToBlockchain(dippe *abe.DIPPE, key []byte) {
	//get access structure from blockchain
	dtObject := core.ReadTDToDTObject("td_sample.json")
	dtObjectByte, err := tools.QueryChaincode(core.ChaincodeId, "GetDTObject", dtObject.Id)
	err = json.Unmarshal(dtObjectByte, &dtObject)
	if err != nil {
		panic(err)
	}
	accessStruct := dtObject.AccessStruct

	bigIntArr := make([]*big.Int, 0)
	accessStructArr := strings.Split(accessStruct, ",")
	for _, s := range accessStructArr {
		int64, _ := strconv.ParseInt(s, 10, 64)
		bigIntArr = append(bigIntArr, big.NewInt(int64))
	}

	policyVec := data.Vector(bigIntArr)

	vecLen := len(accessStructArr)
	// create authorities and their public keys
	auth := make([]*abe.DIPPEAuth, vecLen)
	pubKeys := make([]*abe.DIPPEPubKey, vecLen)
	for i := range auth {
		auth[i], _ = dippe.NewDIPPEAuth(i)
		pubKeys[i] = &auth[i].Pk
	}

	// encrypt the message with the chosen policy give by a policy vector,
	cipher, _ := dippe.Encrypt(string(key), policyVec, pubKeys)

	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	dtObject.CT = CTBytes
	dtObject.IPFSAddress = "IPFS_ADDRESS_123456"
	dtObjectBytes, _ := json.Marshal(dtObject)
	//call update smart contract
	result, err := tools.ExecuteChaincode(core.ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	fmt.Printf("dtObjectBytesSize: %v \n", len(dtObjectBytes))
}
