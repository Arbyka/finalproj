package main

import (
    // "log"
	"os"
	"github.com/joho/godotenv"

    "project-root/config"
    "project-root/route"
)

func main() {
    env := os.Getenv("ENV")
	if env == "docker" {
		_ = godotenv.Load(".env.docker")
	} else {
		_ = godotenv.Load(".env.local")
	}

    config.ConnectDatabase()
    r := route.SetupRouter()
    r.Run(":8080")
}
