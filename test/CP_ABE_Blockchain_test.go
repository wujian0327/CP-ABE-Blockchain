package test

import (
	"CP-ABE-Blockchain/core"
	"CP-ABE-Blockchain/tools"
	"encoding/json"
	"fmt"
	"github.com/fentec-project/gofe/abe"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

const AccessStructure = "((company_A AND repairman) OR (company_B AND consumer) OR admin)"

func TestDeviceCreateDTObject(t *testing.T) {
	core.CreateDTObject("td_sample.json", AccessStructure)
}

func TestDeviceEncryptDataUpload(t *testing.T) {
	core.EncryptDataUploadAndShareKey(tools.RandKey(32))
}

func TestCPABEAttrEncryptt(*testing.T) {
	//Request the AS to obtain PK
	resp, err := http.Get("http://localhost:8888/pkInst")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var pkInst core.PKInst
	err = json.Unmarshal(body, &pkInst)
	if err != nil {
		panic(err)
	}

	key := tools.RandKey(32)
	inst := abe.NewFAME()
	accessStructure := "((A AND B) OR (C AND D) OR E)"

	//CreateDT
	start := time.Now()
	core.CreateDTObject("td_sample.json", accessStructure)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("--------DTModeling time: %s\n", elapsed)

	//Encrypt Data and Upload
	start = time.Now()
	core.EncryptDataUploadAndShareKey(key)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Encrypt Data and Upload time: %s\n", elapsed)

	//Encrypt Key and Save to Blockchain
	start = time.Now()
	_, dtObject := CPABEShareKeyToBlockchain(inst, key)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Encrypt Key and Save to Blockchain time: %s\n", elapsed)

	//Key Generation
	start = time.Now()
	attrsArray := []string{"A", "B"}
	attribKeys := core.KeysGen(attrsArray)
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Key Generation time: %s\n", elapsed)

	//Decryption
	start = time.Now()
	var CT *abe.FAMECipher
	err = json.Unmarshal(dtObject.CT, &CT)
	if err != nil {
		panic(err)
	}

	keyDec, err := pkInst.INST.Decrypt(CT, attribKeys, pkInst.PK)
	if err != nil {
		keyDec = string(key)
	}
	fmt.Println("keyDec: ", keyDec)

	//Decrypt file
	err = tools.DecryptFile("../device_data/data_enc.txt", "../device_data/data_dec.txt", []byte(keyDec))
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("--------Decryption time: %s\n", elapsed)
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
	gamma := []string{"A", "B"}
	keys, _ := pkInst.GenerateAttribKeys(gamma, secKey)
	SKBytes, _ := json.Marshal(keys)
	fmt.Printf("SK Size: %v Bytes\n", len(SKBytes))
}
