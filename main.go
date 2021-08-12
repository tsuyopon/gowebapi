package main

import "gowebapi/server"

func main() {
    r := server.GetRouter()
    r.Run(":8080")
}
