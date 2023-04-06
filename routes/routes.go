package routes

import (
	"github.com/gin-gonic/gin"
)

// AttachRoutes attaches all routes to the gin engine
// this allows for simpler local development and easier testing
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
