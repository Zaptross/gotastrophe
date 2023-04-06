package main

import (
	"github.com/gin-gonic/gin"

	routes "github.com/Zaptross/gotastrophy/routes"
)

func main() {
	r := gin.Default()

	routes.AttachRoutes(r)

	r.Run()
}
