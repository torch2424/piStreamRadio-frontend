package routeHandlers

// Imports
import (
  "fmt"
  "path/filepath"
  "os"
)

import "../renderer"

//3P libs
import (
  "github.com/gin-gonic/gin"
  "github.com/Jeffail/gabs"
  "github.com/dhowden/tag"
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

// Playlist the handler for the "/playlist" route. Will Show Songs on the playlist
func Playlist(c *gin.Context) {
  configJSON, _ := gabs.ParseJSON(renderer.ReadFileAsByte("./config.json"))
  musicDir, _ := configJSON.Path("musicFilesPath").Data().(string)
  filepath.Walk(musicDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
      // Open the file for reading
      file, err := os.Open(path)
      if err != nil {
      	panic(err)
      }

      // Send the file to the meta data tagger
      m, err := tag.ReadFrom(file)
      if err != nil {
      	panic(err)
      }

      // TODO: Add the info to the playlist JSON
      fmt.Println(m.Title())
		}
		return nil
	})

c.Data(200, "text/html; charset=utf-8",
  []byte(renderer.RenderFromPath("./templates/pages/playlist/playlist.html", "./templates/pages/playlist/playlist.json")))
}
