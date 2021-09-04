package util

import (
    "github.com/gin-gonic/gin"
)
type ClientIP struct {
	ClientIP string `json:'clientip"`
}

// 関数名は先頭を大文字にしないとexportできない
func DisplayClientIPAddressAction(c *gin.Context){

//	clientip := ClientIP {
//		ClientIP: "127.0.0.1",
//	}
	clientip := c.ClientIP()

    c.JSON(200, gin.H{"ClientIP": clientip})
}

func DisplayPingAction(c *gin.Context){
	c.JSON(200, gin.H{"message": "pong",})
}
