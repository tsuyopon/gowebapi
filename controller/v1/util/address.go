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

func SomeMethodAction(c *gin.Context){
	var method string
	if c.Request.Method == "GET" {
		method = "get"
	} else if c.Request.Method == "POST" {
		method = "post"
	} else if c.Request.Method == "PUT" {
		method = "put"
	} else if c.Request.Method == "DELETE" {
		method = "delete"
	} else if c.Request.Method == "PATCH" {
		method = "patch"
	} else if c.Request.Method == "HEAD" {
		method = "head"
	} else if c.Request.Method == "OPTIONS" {
		method = "options"
	} else {
		method = "unknown"
	}
	c.JSON(200, gin.H{"message": method})
}

