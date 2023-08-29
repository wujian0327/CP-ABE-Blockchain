package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

type SimpleChaincode struct {
	contractapi.Contract
}

type DTObject struct {
	Id           string      `json:"id"`
	Title        string      `json:"title"`
	Security     string      `json:"security"`
	Properties   interface{} `json:"properties"`
	Actions      interface{} `json:"actions"`
	Events       interface{} `json:"events"`
	AccessStruct string      `json:"access_struct"`
	IPFSAddress  string      `json:"IPFS_address"`
	CT           []byte      `json:"CT"`
}

func (t *SimpleChaincode) PutValue(ctx contractapi.TransactionContextInterface, key string, value string) error {
	err := ctx.GetStub().PutState(key, []byte(value))
	fmt.Printf("put value success,key:%v,value:%v", key, value)
	return err
}

func (t *SimpleChaincode) GetValue(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	b, err := ctx.GetStub().GetState(key)
	if b == nil {
		return "", fmt.Errorf("key doesn't exist")
	}
	return string(b), err
}

func (t *SimpleChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("Init Ledger")
	return nil
}

func (t *SimpleChaincode) Modeling(ctx contractapi.TransactionContextInterface, tdJson string, accessStruct string) (string, error) {
	var dtObject DTObject
	err := json.Unmarshal([]byte(tdJson), &dtObject)
	if err != nil {
		return "", err
	}
	dtObject.AccessStruct = accessStruct
	//save to blockchain
	dtObjectId := "DT_" + dtObject.Id
	dtBytes, err := json.Marshal(dtObject)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(dtObjectId, dtBytes)
	if err != nil {
		return "", err
	}
	fmt.Printf("save DT object to %s , value :%v \n", dtObjectId, dtObject)
	return dtObjectId, nil
}

func (t *SimpleChaincode) Update(ctx contractapi.TransactionContextInterface, tdJson string) (string, error) {
	var dtObject DTObject
	err := json.Unmarshal([]byte(tdJson), &dtObject)
	if err != nil {
		return "", err
	}
	dtObjectId := "DT_" + dtObject.Id
	b, err := ctx.GetStub().GetState(dtObjectId)
	if b == nil {
		return "", fmt.Errorf("key doesn't exist, please use modeling function first")
	}
	dtBytes, err := json.Marshal(dtObject)
	if err != nil {
		return "", err
	}
	err = ctx.GetStub().PutState(dtObjectId, dtBytes)
	if err != nil {
		return "", err
	}
	fmt.Printf("update DT object to %s , value :%v \n", dtObjectId, dtObject)
	return dtObjectId, nil
}

func (t *SimpleChaincode) GetDTObject(ctx contractapi.TransactionContextInterface, dtObjectId string) (string, error) {
	b, err := ctx.GetStub().GetState("DT_" + dtObjectId)
	if b == nil {
		return "", fmt.Errorf("key doesn't exist")
	}
	return string(b), err
}

func (t *SimpleChaincode) SaveIPFSAddress(ctx contractapi.TransactionContextInterface, dtObjectId string, IPFSAddress string) error {
	b, err := ctx.GetStub().GetState(dtObjectId)
	if b == nil {
		return fmt.Errorf("dt object doesn't exist")
	}
	var dtObject DTObject
	err = json.Unmarshal(b, &dtObject)
	if err != nil {
		return err
	}
	dtObject.IPFSAddress = IPFSAddress
	//save to blockchain
	dtBytes, err := json.Marshal(dtObject)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(dtObjectId, dtBytes)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SimpleChaincode{})
	if err != nil {
		log.Panicf("Error creating asset chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting asset chaincode: %v", err)
	}
}
