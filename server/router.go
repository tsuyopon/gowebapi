package server

import (
    "github.com/gin-gonic/gin"

    // バージョンでpackageを階層にする方法がないのでv1のuserかv2のuserを区別するには定義を分離する
    v1_user "gowebapi/controller/v1/user"
    v1_top "gowebapi/controller/v1/top"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
    router := gin.Default()

    router.LoadHTMLGlob("view/**/*.html")  // 全体に適用したい場合(view/common, view/sourceへの適用が必要なため)
//  router.Static("/assets", "./assets")

    // /api/v1グループ分離
    apiv1 := router.Group("/api/v1")
    apiv1.GET("/", v1_top.IndexDisplayAction)
    apiv1.GET("/user", v1_user.UserListDisplayAction)
    apiv1.GET("/user/add", v1_user.UserAddDisplayAction)
    apiv1.GET("/user/edit/:id", v1_user.UserEditDisplayAction)

    return router
}
