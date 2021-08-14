package server

import (
    "github.com/gin-gonic/gin"

    // バージョンでpackageを階層にする方法がないのでv1のuserかv2のuserを区別するには定義を分離する
    v1_user "gowebapi/controller/v1/user"
    v1_top "gowebapi/controller/v1/top"
    "gowebapi/controller/v1/util"
)

func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
    router := gin.Default()

    router.LoadHTMLGlob("view/**/*.html")  // 全体に適用したい場合(view/common, view/sourceへの適用が必要なため)
//  router.Static("/assets", "./assets")

    // /api/v1グループ分離
    apiv1 := router.Group("/api/v1")
    {
        apiv1.GET("/", v1_top.IndexDisplayAction)
        apiv1.GET("/user", v1_user.UserListDisplayAction)
        apiv1.GET("/user/add", v1_user.UserAddDisplayAction)
        apiv1.GET("/user/edit/:id", v1_user.UserEditDisplayAction)

        // util
        apiv1.GET("/clientip", util.DisplayClientIPAddressAction)
        apiv1.GET("/ping", util.DisplayPingAction)
        apiv1.GET("/addHeader", util.DisplayAddHeaderAction)
        apiv1.GET("/printHeader", util.DisplayPrintHeaderAction)
        apiv1.GET("/setCookieHeader", util.SetCookieHeaderAction)
        apiv1.GET("/sleep", util.DisplaySleepAction)
        apiv1.GET("/redirect", util.RedirectAction)
        apiv1.GET("/statusCode/:status", util.StatusCodeAction)

        // 同一関数を指定
        apiv1.GET("/someMethod", util.SomeMethodAction)
        apiv1.POST("/someMethod", util.SomeMethodAction)
        apiv1.PUT("/someMethod", util.SomeMethodAction)
        apiv1.DELETE("/someMethod", util.SomeMethodAction)
        apiv1.PATCH("/someMethod", util.SomeMethodAction)
        apiv1.HEAD("/someMethod", util.SomeMethodAction)
        apiv1.OPTIONS("/someMethod", util.SomeMethodAction)

        apiv1.POST("/printParams", util.PrintParamsAction)
    }

    // Basic認証
    authv1basic := router.Group("/auth/v1", gin.BasicAuth(gin.Accounts{
        "user1": "pass1",
        "user2": "pass2",
    }))
    {
        authv1basic.GET("/secret", util.DisplayPingAction)
    }

    return router
}
