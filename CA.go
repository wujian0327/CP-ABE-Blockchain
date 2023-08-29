package main

import (
	"fmt"
	"github.com/fentec-project/gofe/abe"
	"github.com/gin-gonic/gin"
	"net/http"
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

var keys Keys

func init() {
	inst := abe.NewFAME()

	//generate PK and MSK
	PK, MSK, err := inst.GenerateMasterKeys()
	if err != nil {
		panic(err)
	}
	keys = Keys{
		INST: inst,
		PK:   PK,
		MSK:  MSK,
	}
}

func Setup() (*abe.FAMEPubKey, *abe.FAMESecKey) {
	return keys.PK, keys.MSK
}

func KeysGen(attributes []string) *abe.FAMEAttribKeys {
	println("attributes:", attributes)
	// 解密时构造 属性
	k, err := keys.INST.GenerateAttribKeys(attributes, keys.MSK)
	if err != nil {
		println(err)
	}
	return k
}

func main() {
	r := gin.Default()
	r.GET("/pkInst", func(c *gin.Context) {
		pkInst := PKInst{
			PK:   keys.PK,
			INST: keys.INST,
		}
		c.JSON(http.StatusOK, pkInst)
	})
	r.POST("/KeyGen", func(c *gin.Context) {
		json := make(map[string][]string)
		err := c.BindJSON(&json)
		if err != nil {
			return
		}
		fmt.Printf("%v", &json)
		c.JSON(http.StatusOK, KeysGen(json["attrs"]))
	})
	err := r.Run("0.0.0.0:8888")
	if err != nil {
		return
	}
}
