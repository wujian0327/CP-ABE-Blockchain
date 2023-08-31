package main

import (
	"CP-ABE-Blockchain/core"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// start AS server
	r := gin.Default()
	r.GET("/pkInst", func(c *gin.Context) {
		pkInst := core.PKInst{
			PK:   core.K.PK,
			INST: core.K.INST,
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
		c.JSON(http.StatusOK, core.KeysGen(json["attrs"]))
	})
	err := r.Run("0.0.0.0:8888")
	if err != nil {
		return
	}
}
