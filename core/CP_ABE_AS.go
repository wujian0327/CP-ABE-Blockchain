package core

import (
	"github.com/fentec-project/gofe/abe"
)

type Keys struct {
	INST *abe.FAME
	PK   *abe.FAMEPubKey
	MSK  *abe.FAMESecKey
}

type PKInst struct {
	INST *abe.FAME
	PK   *abe.FAMEPubKey
}

var K Keys

func init() {
	inst := abe.NewFAME()

	//generate PK and MSK
	PK, MSK, err := inst.GenerateMasterKeys()
	if err != nil {
		panic(err)
	}
	K = Keys{
		INST: inst,
		PK:   PK,
		MSK:  MSK,
	}
}

func Setup() (*abe.FAMEPubKey, *abe.FAMESecKey) {
	return K.PK, K.MSK
}

func KeysGen(attributes []string) *abe.FAMEAttribKeys {
	println("attributes:", attributes)
	// 解密时构造 属性
	k, err := K.INST.GenerateAttribKeys(attributes, K.MSK)
	if err != nil {
		println(err)
	}
	return k
}
