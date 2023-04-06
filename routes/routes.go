package routes

import (
	"github.com/gin-gonic/gin"
)

func AttachRoutes(r *gin.Engine) {
	r.GET("/hello-world", GETHelloWorldHandler)
	r.GET("/go-away", GETGoAwayHandler)
}

func GETHelloWorldHandler(c *gin.Context) {
	c.String(200, "Hello World")
}

func GETGoAwayHandler(c *gin.Context) {
	c.String(200, "Go Away")
}
