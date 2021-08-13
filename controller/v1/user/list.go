package user

import (
    "github.com/gin-gonic/gin"
)

func UserListDisplayAction(c *gin.Context){
    c.HTML(200, "user-list.html", gin.H{})
}
