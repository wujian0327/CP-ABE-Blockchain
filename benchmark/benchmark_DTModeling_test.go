package benchmark

import (
	"CP-ABE-Blockchain/core"
	"CP-ABE-Blockchain/tools"
	"encoding/json"
	"fmt"
	"github.com/fentec-project/gofe/abe"
	"testing"
)

const AccessStructure = "((company_A AND repairman) OR (company_B AND consumer) OR admin)"
const DtId = "3deca264-4f90-4321-a5ea-f197e6a1c7cf"
const IPFSAddr = "https://ipfs.io/ipfs/QmRuCN5AAxz9Z1hRBajJKRPfMEYvBdi6xWAoXHeYkSdujm?filename=1.jpeg"

func BenchmarkReadTDFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dtObject := core.ReadTDToDTObject("td_sample.json")
		marshal, _ := json.Marshal(dtObject)
		fmt.Printf("dtObject:%s \n", marshal)
	}
}

//2 org, 3 peer :  0.43s
func BenchmarkModeling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.CreateDTObject("td_sample.json", AccessStructure)
	}
}

func BenchmarkSaveCT(b *testing.B) {
	inst := abe.NewFAME()
	PK, _, _ := inst.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(AccessStructure, false)
	key := tools.RandKey(32)
	cipher, _ := inst.Encrypt(string(key), msp, PK)
	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	fmt.Printf(string(CTBytes), "\n")
	//core.CreateDTObject("td_sample.json", AccessStructure)
	//invokeTimeSum := 0
	//for i := 0; i < b.N; i++ {
	//	invokeTime := core.SaveCT(DtId, CTBytes)
	//	invokeTimeSum += invokeTime
	//}
	//fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//24          42031804 ns/op
//2 org, 3 peer :  0.0433s
//2 org, 5 peer :  0.0458s
//2 org, 7 peer :  0.0470s
//2 org, 9 peer :  0.0493s
//3 org, 3 peer :  0.0579s
//3 org, 5 peer :  0.0611s
//3 org, 7 peer :  0.0679s
//3 org, 9 peer :  0.0722s
//4 org, 3 peer :  0.0609s
//4 org, 5 peer :  0.0631s
//4 org, 7 peer :  0.0688s
//4 org, 9 peer :  0.0743s
//avg invoke time : 31.57 ms
func BenchmarkCreateDTObject_1(b *testing.B) {
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.CreateDTObject("td_files/td_sample_1kb.json", AccessStructure)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

// 30          44847350 ns/op
//avg invoke time : 32.21 ms
func BenchmarkCreateDTObject_2(b *testing.B) {
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.CreateDTObject("td_files/td_sample_2kb.json", AccessStructure)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//26          43668515 ns/op
//avg invoke time : 33.70 ms
func BenchmarkCreateDTObject_4(b *testing.B) {
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.CreateDTObject("td_files/td_sample_4kb.json", AccessStructure)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//21          48832114 ns/op
//avg invoke time : 33.28 ms
func BenchmarkCreateDTObject_8(b *testing.B) {
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.CreateDTObject("td_files/td_sample_8kb.json", AccessStructure)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//21          55205638 ns/op
// avg invoke time : 42.5 ms
func BenchmarkCreateDTObject_16(b *testing.B) {
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.CreateDTObject("td_files/td_sample_16kb.json", AccessStructure)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//28          39974454 ns/op
//avg invoke time : 29.23 ms
func BenchmarkSaveIPFS_1(b *testing.B) {
	core.CreateDTObject("td_files/td_sample_1kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveIPFSAddr(DtId, IPFSAddr)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//31          43009177 ns/op
//avg invoke time : 31.83 ms
func BenchmarkSaveIPFS_2(b *testing.B) {
	core.CreateDTObject("td_files/td_sample_2kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveIPFSAddr(DtId, IPFSAddr)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//24          45741542 ns/op
//avg invoke time : 33.90 ms
func BenchmarkSaveIPFS_4(b *testing.B) {
	core.CreateDTObject("td_files/td_sample_4kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveIPFSAddr(DtId, IPFSAddr)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

// 30          39267257 ns/op
//avg invoke time : 32.03 ms
func BenchmarkSaveIPFS_8(b *testing.B) {
	core.CreateDTObject("td_files/td_sample_8kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveIPFSAddr(DtId, IPFSAddr)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

// 24          45322004 ns/op
//avg invoke time : 40.00 ms
func BenchmarkSaveIPFS_16(b *testing.B) {
	core.CreateDTObject("td_files/td_sample_16kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveIPFSAddr(DtId, IPFSAddr)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//25          48771872 ns/op
//avg invoke time : 33.67 ms
func BenchmarkSaveCT_1(b *testing.B) {
	inst := abe.NewFAME()
	PK, _, _ := inst.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(AccessStructure, false)
	key := tools.RandKey(32)
	cipher, _ := inst.Encrypt(string(key), msp, PK)
	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	core.CreateDTObject("td_files/td_sample_1kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveCT(DtId, CTBytes)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

// 25          51720880 ns/op
//avg invoke time : 34.4 ms
func BenchmarkSaveCT_2(b *testing.B) {
	inst := abe.NewFAME()
	PK, _, _ := inst.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(AccessStructure, false)
	key := tools.RandKey(32)
	cipher, _ := inst.Encrypt(string(key), msp, PK)
	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	core.CreateDTObject("td_files/td_sample_2kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveCT(DtId, CTBytes)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//22          55730632 ns/op
//avg invoke time : 42.00 ms
func BenchmarkSaveCT_4(b *testing.B) {
	inst := abe.NewFAME()
	PK, _, _ := inst.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(AccessStructure, false)
	key := tools.RandKey(32)
	cipher, _ := inst.Encrypt(string(key), msp, PK)
	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	core.CreateDTObject("td_files/td_sample_4kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveCT(DtId, CTBytes)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//22          55730632 ns/op
//avg invoke time : 45.10 ms
func BenchmarkSaveCT_8(b *testing.B) {
	inst := abe.NewFAME()
	PK, _, _ := inst.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(AccessStructure, false)
	key := tools.RandKey(32)
	cipher, _ := inst.Encrypt(string(key), msp, PK)
	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	core.CreateDTObject("td_files/td_sample_8kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveCT(DtId, CTBytes)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}

//28          59688654 ns/op
//avg invoke time : 45.29 ms
func BenchmarkSaveCT_16(b *testing.B) {
	inst := abe.NewFAME()
	PK, _, _ := inst.GenerateMasterKeys()
	msp, _ := abe.BooleanToMSP(AccessStructure, false)
	key := tools.RandKey(32)
	cipher, _ := inst.Encrypt(string(key), msp, PK)
	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	core.CreateDTObject("td_files/td_sample_16kb.json", AccessStructure)
	invokeTimeSum := 0
	for i := 0; i < b.N; i++ {
		invokeTime := core.SaveCT(DtId, CTBytes)
		invokeTimeSum += invokeTime
	}
	fmt.Printf("avg invoke time : %v ms \n", float64(invokeTimeSum)/float64(b.N))
}
