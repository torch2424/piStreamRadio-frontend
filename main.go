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
	raymond.RegisterPartial("boilerplate-style", string(renderer.ReadFileAsByte("./templates/partials/boilerplate.style.html")))
	raymond.RegisterPartial("style", string(renderer.ReadFileAsByte("./templates/partials/style.html")))
	raymond.RegisterPartial("head", string(renderer.ReadFileAsByte("./templates/partials/head.html")))
	raymond.RegisterPartial("footer", string(renderer.ReadFileAsByte("./templates/partials/footer.html")))

	// Define our Routes
	r.GET("/", routeHandlers.Home)
	r.GET("/about", routeHandlers.About)
	r.GET("/faq", routeHandlers.FAQ)
	r.GET("/contact", routeHandlers.Contact)
	r.GET("/playlist", routeHandlers.Playlist)

	// TODO: Handle Password Hashing
	// bcrypt still great in 2017
	// https://gowebexamples.com/password-hashing/

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
