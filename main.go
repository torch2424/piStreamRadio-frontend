package main

//Our libs
import "./routeHandlers"

//3P libs
import "github.com/gin-gonic/gin"

func main() {
	// Create our Router with Gin
	r := gin.Default()

	// Define our Routes
	r.GET("/", routeHandlers.Home)

	// TODO: Remove this example route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
