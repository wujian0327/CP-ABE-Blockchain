package core

import (
	"github.com/fentec-project/gofe/abe"
)

type MAKeys struct {
	INST *abe.MAABE
	PK   []*abe.MAABEPubKey
}

var MaPK MAKeys

func init() {
	maabe := abe.NewMAABE()

	// create three authorities, each with two attributes
	attribs1 := []string{"auth1:at1", "auth1:at2"}
	attribs2 := []string{"auth2:at1", "auth2:at2"}
	attribs3 := []string{"auth3:at1", "auth3:at2"}
	auth1, _ := maabe.NewMAABEAuth("auth1", attribs1)
	auth2, _ := maabe.NewMAABEAuth("auth2", attribs2)
	auth3, _ := maabe.NewMAABEAuth("auth3", attribs3)
	// define the set of all public keys we use
	pks := []*abe.MAABEPubKey{auth1.PubKeys(), auth2.PubKeys(), auth3.PubKeys()}

	MaPK = MAKeys{
		INST: maabe,
		PK:   pks,
	}
}

func KpKeysGen(attributes []string) *abe.FAMEAttribKeys {
	println("attributes:", attributes)
	// 解密时构造 属性
	k, err := K.INST.GenerateAttribKeys(attributes, K.MSK)
	if err != nil {
		println(err)
	}
	return k
}
