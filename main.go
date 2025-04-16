package main

import (
    "project-root/config"
    "project-root/route"
)

func main() {
    config.ConnectDatabase()
    r := route.SetupRouter()
    r.Run(":8080")
}
