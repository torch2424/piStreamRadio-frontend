package routeHandlers

// Imports
import (
  "fmt"
)

import "../renderer"

//3P libs
import (
  "github.com/gin-gonic/gin"
  "github.com/Jeffail/gabs"
)

// Home The handler for the "/" route. Will simply explain the page and other routes
func Home(c *gin.Context) {

  // Digest JSON in gabs to allow editing
  jsonParsed, _ := gabs.ParseJSON(renderer.ReadFileAsByte("./templates/pages/home/home.json"))
  jsonParsed.SetP("I changed the body again with gabssss! lmaooo ayyeee", "body");

  fmt.Println("Rendering / ...")

  // Return our html
  c.Data(200, "text/html; charset=utf-8",
    []byte(renderer.RenderFromRawJSON("./templates/pages/home/home.html",
      []byte(jsonParsed.String()))))
}
