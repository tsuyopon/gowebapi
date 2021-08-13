package util

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func DisplayAddHeaderAction(c *gin.Context){
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    c.Header("Access-Control-Max-Age", "86400")
    c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    c.String(200, "Add CORS Headers")
}

func DisplayPrintHeaderAction(c *gin.Context){
    var str string
    for k, v := range c.Request.Header {
      str += fmt.Sprint(k,":",v,"\n")
    }
    c.String(200, str)
}

func SetCookieHeaderAction(c *gin.Context){
	c.SetCookie("user", "hoge", 3600, "/", "localhost", false, true)
}
