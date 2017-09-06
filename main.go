package main

//Our libs
import "./routeHandlers"
import "./renderer"

//3P libs
import (
	"github.com/gin-gonic/gin"
	"github.com/aymerick/raymond"
)

func main() {
	// Create our Router with Gin
	r := gin.Default()

	// Register all of our partials with raymond
	raymond.RegisterPartial("head", string(renderer.ReadFileAsByte("./templates/partials/head.html")))
	raymond.RegisterPartial("footer", string(renderer.ReadFileAsByte("./templates/partials/footer.html")))

	// Define our Routes
	r.GET("/", routeHandlers.Home)
	r.GET("/playlist", routeHandlers.Playlist)

	// TODO: Handle Password Hashing
	// bcrypt still great in 2017
	// https://gowebexamples.com/password-hashing/

	// TODO: Remove this example route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
