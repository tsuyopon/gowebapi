package user

import (
    "github.com/gin-gonic/gin"
)

func UserAddDisplayAction(c *gin.Context){
    c.HTML(200, "user-add.html", gin.H{})
}
