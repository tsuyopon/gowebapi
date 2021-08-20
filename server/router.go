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

    // 静的なassetsを配置する場所
    router.Static("/images", "./assets/images")
    router.Static("/css", "./assets/css")
    router.Static("/js", "./assets/js")
    router.Static("/html", "./assets/html")

    // web/v1グループ
    webv1:= router.Group("/web/v1")
    {
        webv1.GET("/", v1_top.IndexDisplayAction)
        webv1.GET("/user", v1_user.UserListDisplayAction)
        webv1.GET("/user/add", v1_user.UserAddDisplayAction)
        webv1.GET("/user/edit/:id", v1_user.UserEditDisplayAction)
    }

    // /api/v1グループ分離
    apiv1 := router.Group("/api/v1")
    {
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
