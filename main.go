package main

import (
    "io"
    "os"
    "github.com/gin-gonic/gin"
    "gowebapi/server"
)

func main() {

    // Disable Console Color, you don't need console color when writing the logs to file.
    gin.DisableConsoleColor()

    // Logging to a file.
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    // Use the following code if you need to write the logs to file and console at the same time.
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    r := server.GetRouter()

    // どちらか片方のみ有効にしてください。両方は適用されません。
//    r.Run(":8080")
    r.RunTLS(":4433", "./conf/testdata/server.crt", "./conf/testdata/server.key")

}
