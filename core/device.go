package core

import (
	"CP-ABE-Blockchain/tools"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fentec-project/gofe/abe"
	"io/ioutil"
	"net/http"
	"os"
)

const ChaincodeId = "DTModeling"

type DTObject struct {
	Id           string      `json:"id"`
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	Security     string      `json:"security"`
	Properties   interface{} `json:"properties"`
	Actions      interface{} `json:"actions"`
	Events       interface{} `json:"events"`
	AccessStruct string      `json:"access_struct"`
	IPFSAddress  string      `json:"IPFS_address"`
	CT           []byte      `json:"CT"`
}

type Attributes struct {
	Attrs []string `json:"attrs"`
}

func CreateDTObject(tdFilePath string, accessStruct string) int {
	dtObject := ReadTDToDTObject(tdFilePath)
	dtObject.AccessStruct = accessStruct
	dtObjectBytes, err := json.Marshal(dtObject)
	if err != nil {
		panic(err)
	}
	//call smart contract
	result, invokeTime, err := tools.ExecuteChaincode(ChaincodeId, "Modeling", string(dtObjectBytes), accessStruct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	return int(invokeTime)
}

func UpdateDTObject(tdFilePath string) {
	dtObject := ReadTDToDTObject(tdFilePath)
	dtObjectBytes, err := json.Marshal(dtObject)
	if err != nil {
		panic(err)
	}
	//call smart contract
	result, _, err := tools.ExecuteChaincode(ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
}

func SaveIPFSAddr(dtObjectId string, IPFSAddress string) int {
	//call smart contract
	result, invokeTime, err := tools.ExecuteChaincode(ChaincodeId, "SaveIPFSAddress", dtObjectId, IPFSAddress)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	return int(invokeTime)
}

func SaveCT(dtObjectId string, CT []byte) int {
	//call smart contract
	result, invokeTime, err := tools.ExecuteChaincode(ChaincodeId, "SaveCT", dtObjectId, string(CT))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	return int(invokeTime)
}

func EncryptDataUploadAndShareKey(key []byte) {
	//Randomly generate AES keys
	fmt.Println("key: ", string(key))
	//encrypt file
	err := tools.EncryptFile("../device_data/data.txt", "../device_data/data_enc.txt", key)
	if err != nil {
		panic(err)
	}

	//storage to ipfs
	filePath := "device_data/data_enc.txt"
	ipfsAddress := shareFileToIpfs(filePath)
	ShareKeyToBlockchain(key, ipfsAddress)
}

func shareFileToIpfs(path string) string {
	uri := "http://127.0.0.1:5001/api/v0/add"
	byte, err := ioutil.ReadFile(path)
	res, err := http.Post(uri, "multipart/form-data", bytes.NewReader(byte))
	if err != nil {
		fmt.Println("err=", err)
	}
	//http返回的response的body必须close,否则就会有内存泄露
	defer func() {
		res.Body.Close()
		fmt.Println("finish")
	}()
	//读取body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(" post err=", err)
	}
	fmt.Println(string(body))
	return string(body)
}

func ShareKeyToBlockchain(key []byte, ipfsAddress string) {
	//Request the AS to obtain PK
	resp, err := http.Get("http://localhost:8888/pkInst")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var pkInst PKInst
	err = json.Unmarshal(body, &pkInst)
	if err != nil {
		panic(err)
	}

	//get access structure from blockchain
	dtObject := ReadTDToDTObject("td_sample.json")
	dtObjectByte, err := tools.QueryChaincode(ChaincodeId, "GetDTObject", dtObject.Id, "")
	err = json.Unmarshal(dtObjectByte, &dtObject)
	if err != nil {
		panic(err)
	}
	accessStruct := dtObject.AccessStruct
	msp, err := abe.BooleanToMSP(accessStruct, false)
	if err != nil {
		panic(err)
	}
	cipher, err := pkInst.INST.Encrypt(string(key), msp, pkInst.PK)

	//save CP and ipfsAddress to blockchain
	CTBytes, _ := json.Marshal(cipher)
	dtObject.CT = CTBytes
	dtObject.IPFSAddress = ipfsAddress
	dtObjectBytes, _ := json.Marshal(dtObject)
	//call update smart contract
	result, _, err := tools.ExecuteChaincode(ChaincodeId, "Update", string(dtObjectBytes))
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
	fmt.Printf("dtObjectBytesSize: %v \n", len(dtObjectBytes))
}

func ReadTDToDTObject(tdFilePath string) DTObject {
	//read tdSample file
	content, err := os.ReadFile(tdFilePath)
	if err != nil {
		panic(err)
	}
	var dtObject DTObject
	err = json.Unmarshal(content, &dtObject)
	if err != nil {
		panic(err)
	}
	return dtObject
}
