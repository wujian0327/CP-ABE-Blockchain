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
	"time"
)

var gamma []string

func init() {
	gamma = []string{"A", "B", "C", "D", "E"}
}

//CPABE Attr5 test
//12          97299950 ns/op  9.70 TX/s
//dtObjectBytesSize: 10436
//encrypt time : 23 ms
//avg invoke time : 34 ms
//Ciphertex Size: 7342 Bytes
//SK Size: 4360 Bytes
//msp Size: 96 Bytes
func BenchmarkCPABEAttrEncrypt_5(b *testing.B) {
	key := tools.RandKey(32)
	inst := abe.NewFAME()
	accessStructure := "((A AND B) OR (C AND D) OR E)"
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime, _ := CPABEShareKeyToBlockchain(inst, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//CPABE Attr10 test
//8         126052138 ns/op  7.93 TX/s
//dtObjectBytesSize: 16295
//encrypt time : 56 ms
//avg invoke time : 39 ms
//Ciphertex Size: 11711 Bytes
//SK Size: 4365 Bytes
//msp Size: 198 Bytes
func BenchmarkCPABEAttrEncrypt_10(b *testing.B) {
	key := tools.RandKey(32)
	inst := abe.NewFAME()
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J)"
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime, _ := CPABEShareKeyToBlockchain(inst, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//CPABE Attr15 test
//6         179794433 ns/op  5.58
//dtObjectBytesSize: 22194
//encrypt time : 102 ms
//avg invoke time : 48 ms
//Ciphertex Size: 16114 Bytes
//SK Size: 4351 Bytes
//msp Size: 340 Bytes
func BenchmarkCPABEAttrEncrypt_15(b *testing.B) {
	key := tools.RandKey(32)
	inst := abe.NewFAME()
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O)"
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime, _ := CPABEShareKeyToBlockchain(inst, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//CPABE Attr20 test
//4         289363150 ns/op  3.46 TX/s
//dtObjectBytesSize: 28141
//encrypt time : 171 ms
//avg invoke time : 57 ms
//Ciphertex Size: 20586 Bytes
//SK Size: 4352 Bytes
//msp Size: 522 Bytes
func BenchmarkCPABEAttrEncrypt_20(b *testing.B) {
	key := tools.RandKey(32)
	inst := abe.NewFAME()
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O" +
		" OR (P AND Q) OR (R AND S) OR T)"
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime, _ := CPABEShareKeyToBlockchain(inst, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//CPABE Attr25 test
//3         342064033 ns/op  2.92
//dtObjectBytesSize: 34168
//encrypt time : 251 ms
//avg invoke time : 68 ms
//Ciphertex Size: 25038 Bytes
//SK Size: 4360 Bytes
//msp Size: 744 Bytes
func BenchmarkCPABEAttrEncrypt_25(b *testing.B) {
	key := tools.RandKey(32)
	inst := abe.NewFAME()
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O" +
		" OR (P AND Q) OR (R AND S) OR T OR (U AND V) OR (W AND X) OR Y)"
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime, _ := CPABEShareKeyToBlockchain(inst, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//CPABE Attr30 test
// 3         468285567 ns/op   2.13 TX/s
//dtObjectBytesSize: 40236
//encrypt time : 349 ms
//avg invoke time : 76 ms
//Ciphertex Size: 29560 Bytes
//SK Size: 4355 Bytes
//msp Size: 1010 Bytes
func BenchmarkCPABEAttrEncrypt_30(b *testing.B) {
	key := tools.RandKey(32)
	inst := abe.NewFAME()
	accessStructure := "((A AND B) OR (C AND D) OR E OR (F AND G) OR (H AND I) OR J OR (K AND L) OR (M AND N) OR O" +
		" OR (P AND Q) OR (R AND S) OR T OR (U AND V) OR (W AND X) OR Y" +
		" OR (Z AND AA) OR (AB AND AC) OR AD )"
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime, _ := CPABEShareKeyToBlockchain(inst, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//MAABE Attr5
//9         124503944 ns/op  8.06 TX/s
//dtObjectBytesSize: 17878
//encrypt time : 12 ms
//avg invoke time : 33 ms
//Ciphertex Size: 12886 Bytes
//SK Size: 980 Bytes
//msp Size: 136 Bytes
func BenchmarkMAABEAttr_5(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth2:at1) OR (auth1:at2 AND auth2:at2)) OR (auth3:at1)"
	maabe := abe.NewMAABE()
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := MAABEShareKeyToBlockchain(maabe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)

}

//MAABE Attr10
//7         176703929 ns/op  5.68 TX/s
//dtObjectBytesSize: 33519
//encrypt time : 24 ms
//avg invoke time : 50 ms
//Ciphertex Size: 24564 Bytes
//SK Size: 979 Bytes
//msp Size: 278 Bytes
func BenchmarkMAABEAttr_10(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5))"
	maabe := abe.NewMAABE()
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := MAABEShareKeyToBlockchain(maabe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)

}

//MAABE Attr15
//6         201211267 ns/op    4.97 TX/s
//dtObjectBytesSize: 49208
//encrypt time : 35 ms
//avg invoke time : 69 ms
//Ciphertex Size: 36238 Bytes
//SK Size: 979 Bytes
//msp Size: 460 Bytes
func BenchmarkMAABEAttr_15(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5))"
	maabe := abe.NewMAABE()
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := MAABEShareKeyToBlockchain(maabe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)

}

//MAABE Attr20
//6         207949617 ns/op    4.83TX/s
//encrypt time : 46 ms
//dtObjectBytesSize: 64902
//encrypt time : 46 ms
//avg invoke time : 70 ms
//Ciphertex Size: 47998 Bytes
//SK Size: 979 Bytes
//msp Size: 683 Bytes
func BenchmarkMAABEAttr_20(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5)" +
		" OR (auth1:at6 AND auth1:at7) OR (auth1:at8 AND auth1:at9) OR (auth1:at10))"
	maabe := abe.NewMAABE()
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := MAABEShareKeyToBlockchain(maabe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//MAABE Attr25
//5         278099380 ns/op   3.59TX/s
//dtObjectBytesSize: 80704
//encrypt time : 55 ms
//avg invoke time : 80 ms
//Ciphertex Size: 59798 Bytes
//SK Size: 979 Bytes
//msp Size: 946 Bytes
func BenchmarkMAABEAttr_25(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5)" +
		" OR (auth2:at6 AND auth2:at7) OR (auth2:at8 AND auth2:at9) OR (auth2:at10)" +
		" OR (auth1:at6 AND auth1:at7) OR (auth1:at8 AND auth1:at9) OR (auth1:at10))"
	maabe := abe.NewMAABE()
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := MAABEShareKeyToBlockchain(maabe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//MAABE Attr30
//4         308813025 ns/op  3.24TX/s
//dtObjectBytesSize: 96498
//encrypt time : 65 ms
//avg invoke time : 86 ms
//Ciphertex Size: 71598 Bytes
//SK Size: 972 Bytes
//msp Size: 1249 Bytes
func BenchmarkMAABEAttr_30(b *testing.B) {
	key := tools.RandKey(32)
	accessStructure := "((auth1:at1 AND auth1:at2) OR (auth1:at3 AND auth1:at4) OR (auth1:at5)" +
		" OR (auth2:at1 AND auth2:at2) OR (auth2:at3 AND auth2:at4) OR (auth2:at5)" +
		" OR (auth3:at1 AND auth3:at2) OR (auth3:at3 AND auth3:at4) OR (auth3:at5)" +
		" OR (auth3:at6 AND auth3:at7) OR (auth3:at8 AND auth3:at9) OR (auth3:at10)" +
		" OR (auth2:at6 AND auth2:at7) OR (auth2:at8 AND auth2:at9) OR (auth2:at10)" +
		" OR (auth1:at6 AND auth1:at7) OR (auth1:at8 AND auth1:at9) OR (auth1:at10))"
	maabe := abe.NewMAABE()
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := MAABEShareKeyToBlockchain(maabe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//DIPPE Attr5
// 8         130792738 ns/op  7.69 TX/s
//dtObjectBytesSize: 11385
//encrypt time : 61 ms
//avg invoke time : 36 ms
//Ciphertex Size: 8077 Bytes
//SK Size: 11667 Bytes
//msp Size: 11 Bytes
func BenchmarkDIPPEAttr_5(b *testing.B) {
	key := tools.RandKey(32)
	// choose a policy vector
	accessStructure := "1,2,3,4,5"
	dippe, _ := abe.NewDIPPE(3)
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := DIPPEShareKeyToBlockchain(dippe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//DIPPE Attr10
//6         214917867 ns/op  4.67 TX/s
//dtObjectBytesSize: 18967
//encrypt time : 112 ms
//avg invoke time : 52 ms
//Ciphertex Size: 13747 Bytes
//SK Size: 11689 Bytes
//msp Size: 22 Bytes
func BenchmarkDIPPEAttr_10(b *testing.B) {
	key := tools.RandKey(32)
	// choose a policy vector
	accessStructure := "1,2,3,4,5,6,7,8,9,10"
	dippe, _ := abe.NewDIPPE(3)
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := DIPPEShareKeyToBlockchain(dippe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//DIPPE Attr15
//4         279850600 ns/op   3.58TX/s
//dtObjectBytesSize: 26582
//encrypt time : 181 ms
//avg invoke time : 53 ms
//Ciphertex Size: 19441 Bytes
//SK Size: 11733 Bytes
//msp Size: 37 Bytes
func BenchmarkDIPPEAttr_15(b *testing.B) {
	key := tools.RandKey(32)
	// choose a policy vector
	accessStructure := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15"
	dippe, _ := abe.NewDIPPE(3)
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := DIPPEShareKeyToBlockchain(dippe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//DIPPE Attr20
//4         310063675 ns/op  3.22 TX/s
//dtObjectBytesSize: 34201
//encrypt time : 232 ms
//avg invoke time : 46 ms
//Ciphertex Size: 25134 Bytes
//SK Size: 11745 Bytes
//msp Size: 52 Bytes
func BenchmarkDIPPEAttr_20(b *testing.B) {
	key := tools.RandKey(32)
	// choose a policy vector
	accessStructure := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20"
	dippe, _ := abe.NewDIPPE(3)
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := DIPPEShareKeyToBlockchain(dippe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//DIPPE Attr25
//3         398355400 ns/op   2.51TX/s
//dtObjectBytesSize: 41796
//encrypt time : 286 ms
//avg invoke time : 74 ms
//Ciphertex Size: 30799 Bytes
//SK Size: 11782 Bytes
//msp Size: 67 Bytes
func BenchmarkDIPPEAttr_25(b *testing.B) {
	key := tools.RandKey(32)
	// choose a policy vector
	accessStructure := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20" +
		",21,22,23,24,25"
	dippe, _ := abe.NewDIPPE(3)
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := DIPPEShareKeyToBlockchain(dippe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

//DIPPE Attr30
//3         495481867 ns/op   2.02TX/s
//dtObjectBytesSize: 49395
//encrypt time : 346  ms
//avg invoke time : 85 ms
//Ciphertex Size: 36526 Bytes
//SK Size: 11797 Bytes
//msp Size: 82 Bytes
func BenchmarkDIPPEAttr_30(b *testing.B) {
	key := tools.RandKey(32)
	// choose a policy vector
	accessStructure := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20" +
		",21,22,23,24,25,26,27,28,29,30"
	dippe, _ := abe.NewDIPPE(3)
	core.CreateDTObject("td_sample.json", accessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := DIPPEShareKeyToBlockchain(dippe, key)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", invokeTimeSum/b.N)
}

func CPABEShareKeyToBlockchain(pkInst *abe.FAME, key []byte) (int, core.DTObject) {
	//get access structure from blockchain
	dtObject := core.ReadTDToDTObject("td_sample.json")
	dtObjectByte, err := tools.QueryChaincode(core.ChaincodeId, "GetDTObject", dtObject.Id, "")
	err = json.Unmarshal(dtObjectByte, &dtObject)
	if err != nil {
		panic(err)
	}
	accessStruct := dtObject.AccessStruct
	CPABEEncrypt(pkInst, key, accessStruct, &dtObject)
	dtObjectBytes, _ := json.Marshal(dtObject)
	//call update smart contract
	_, invokeTime, err := tools.ExecuteChaincode(core.ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	//fmt.Printf("result: %s \n", result)
	fmt.Printf("dtObjectBytesSize: %v \n", len(dtObjectBytes))
	return int(invokeTime), dtObject
}

func CPABEEncrypt(pkInst *abe.FAME, key []byte, accessStruct string, dtObject *core.DTObject) {
	start := time.Now().UnixMilli()
	PK, secKey, err := pkInst.GenerateMasterKeys()
	msp, err := abe.BooleanToMSP(accessStruct, false)
	mspBytes, _ := json.Marshal(msp)
	fmt.Printf("msp Size: %v Bytes\n", len(mspBytes))
	if err != nil {
		panic(err)
	}
	cipher, err := pkInst.Encrypt(string(key), msp, PK)
	end := time.Now().UnixMilli()
	//save CT and ipfsAddress
	CTBytes, _ := json.Marshal(cipher)
	dtObject.CT = CTBytes
	dtObject.IPFSAddress = "IPFS_ADDRESS_123456"
	fmt.Printf("encrypt time : %v ms\n", end-start)
	fmt.Printf("Ciphertex Size: %v Bytes\n", len(CTBytes))

	keys, _ := pkInst.GenerateAttribKeys(gamma, secKey)
	SKBytes, _ := json.Marshal(keys)
	fmt.Printf("SK Size: %v Bytes\n", len(SKBytes))
}

func MAABEShareKeyToBlockchain(maabe *abe.MAABE, key []byte) int {
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
	dtObjectByte, _ := tools.QueryChaincode(core.ChaincodeId, "GetDTObject", dtObject.Id, "")
	_ = json.Unmarshal(dtObjectByte, &dtObject)

	accessStructure := dtObject.AccessStruct
	// create a msp struct out of the boolean formula
	start := time.Now().UnixMilli()
	msp, err := abe.BooleanToMSP(accessStructure, false)
	if err != nil {
		panic(err)
	}
	mspBytes, _ := json.Marshal(msp)
	fmt.Printf("msp Size: %v Bytes\n", len(mspBytes))
	// define the set of all public keys we use
	pks := []*abe.MAABEPubKey{auth1.PubKeys(), auth2.PubKeys(), auth3.PubKeys()}

	// encrypt the message with the decryption policy in msp
	ct, _ := maabe.Encrypt(string(key), msp, pks)
	end := time.Now().UnixMilli()
	fmt.Printf("encrypt time : %v ms\n", end-start)

	//save CP and ipfsAddress to blockchain

	CTBytes, _ := json.Marshal(ct)
	fmt.Printf("Ciphertex Size: %v Bytes\n", len(CTBytes))
	dtObject.CT = CTBytes
	dtObject.IPFSAddress = "IPFS_ADDRESS_123456"
	dtObjectBytes, _ := json.Marshal(dtObject)
	//call update smart contract
	result, invokeTime, err := tools.ExecuteChaincode(core.ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	fmt.Printf("dtObjectBytesSize: %v \n", len(dtObjectBytes))

	// authority 1 issues keys to user
	gid := "gid1"
	keys1, err := auth1.GenerateAttribKeys(gid, attribs1)
	key11, _ := keys1[0], keys1[1]
	keys2, err := auth2.GenerateAttribKeys(gid, attribs2)
	key21, _ := keys2[0], keys2[1]
	keys3, err := auth3.GenerateAttribKeys(gid, attribs3)
	key31, _ := keys3[0], keys3[1]
	ks1 := []*abe.MAABEKey{key11, key21, key31}
	SKBytes, _ := json.Marshal(ks1)
	fmt.Printf("SK Size: %v Bytes\n", len(SKBytes))
	return int(invokeTime)
}

func DIPPEShareKeyToBlockchain(dippe *abe.DIPPE, key []byte) int {
	//get access structure from blockchain
	dtObject := core.ReadTDToDTObject("td_sample.json")
	dtObjectByte, err := tools.QueryChaincode(core.ChaincodeId, "GetDTObject", dtObject.Id, "")
	err = json.Unmarshal(dtObjectByte, &dtObject)
	if err != nil {
		panic(err)
	}
	accessStruct := dtObject.AccessStruct

	bigIntArr := make([]*big.Int, 0)
	accessStructArr := strings.Split(accessStruct, ",")
	for _, s := range accessStructArr {
		i, _ := strconv.ParseInt(s, 10, 64)
		bigIntArr = append(bigIntArr, big.NewInt(i))
	}

	start := time.Now().UnixMilli()
	policyVec := data.Vector(bigIntArr)
	mspBytes, _ := json.Marshal(policyVec)
	fmt.Printf("msp Size: %v Bytes\n", len(mspBytes))
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
	end := time.Now().UnixMilli()
	fmt.Printf("encrypt time : %v ms\n", end-start)
	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	fmt.Printf("Ciphertex Size: %v Bytes\n", len(CTBytes))
	dtObject.CT = CTBytes
	dtObject.IPFSAddress = "IPFS_ADDRESS_123456"
	dtObjectBytes, _ := json.Marshal(dtObject)
	//call update smart contract
	result, invokeTime, err := tools.ExecuteChaincode(core.ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	fmt.Printf("dtObjectBytesSize: %v \n", len(dtObjectBytes))

	userGID := "someGID"
	userVec := data.Vector([]*big.Int{big.NewInt(0), big.NewInt(1),
		big.NewInt(1), big.NewInt(-3), big.NewInt(6)})

	// authorities generate decryption keys for the user
	userKeys := make([]data.VectorG2, vecLen)
	for i := range userVec {
		userKeys[i], _ = auth[i].DeriveKeyShare(userVec, pubKeys, userGID)
	}
	SKBytes, _ := json.Marshal(userKeys)
	fmt.Printf("SK Size: %v Bytes\n", len(SKBytes))
	return int(invokeTime)
}
