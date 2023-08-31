package core

import (
	"CP-ABE-Blockchain/tools"
	"encoding/json"
	"fmt"
	"github.com/fentec-project/gofe/abe"
	"io/ioutil"
	"net/http"
	"strings"
)

func KeyGen(attrsArray []string) abe.FAMEAttribKeys {
	// 解密时构造 属性
	attrs := Attributes{
		Attrs: attrsArray,
	}
	attrsJson, _ := json.Marshal(attrs)
	targetUrl := "http://localhost:8888/KeyGen"
	payload := strings.NewReader(string(attrsJson))
	//post to get attribKeys
	req, _ := http.NewRequest("POST", targetUrl, payload)
	req.Header.Add("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	var attribKeys abe.FAMEAttribKeys
	err = json.Unmarshal(body, &attribKeys)
	if err != nil {
		panic(err)
	}
	return attribKeys
}

func DecryptData(dtObjectId string, fileSourcePath string, fileTargetPath string) {
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

	//kegGen
	attrsArray := []string{"A", "B"}
	attribKeys := KeyGen(attrsArray)

	//get CT from blockchain
	var dtObject DTObject
	dtObjectByte, err := tools.QueryChaincode(ChaincodeId, "GetDTObject", dtObjectId)
	err = json.Unmarshal(dtObjectByte, &dtObject)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dtObject: %#v \n", dtObject)

	//get AES key
	var CT *abe.FAMECipher
	err = json.Unmarshal(dtObject.CT, &CT)
	if err != nil {
		panic(err)
	}

	key, err := pkInst.INST.Decrypt(CT, &attribKeys, pkInst.PK)
	if err != nil {
		panic(err)
	}
	fmt.Println("key: ", key)

	//Decrypt file
	err = tools.DecryptFile(fileSourcePath, fileTargetPath, []byte(key))
	if err != nil {
		return
	}
	fmt.Println("end")
}
