package routeHandlers

// Imports
import (
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

  // Return our html
  c.Data(200, "text/html; charset=utf-8",
    []byte(renderer.RenderFromRawJSON("./templates/pages/home/home.html",
      []byte(jsonParsed.String()))))
}

// About Handler for the "/about" Route. Simply renders HTML and JSON
func About(c *gin.Context) {
  // Return our html
  c.Data(200, "text/html; charset=utf-8",
    []byte(renderer.RenderFromPath("./templates/pages/about/about.html", "./templates/pages/about/about.json")))
}

// Contact Handler for the "/contact" Route. Simply renders HTML and JSON
func Contact(c *gin.Context) {
  // Return our html
  c.Data(200, "text/html; charset=utf-8",
    []byte(renderer.RenderFromPath("./templates/pages/contact/contact.html", "./templates/pages/contact/contact.json")))
}

// Playlist the handler for the "/playlist" route. Will Show Songs on the playlist
func Playlist(c *gin.Context) {
  // Read our config JSON for the music directory
  configJSON, configJSONErr := gabs.ParseJSON(renderer.ReadFileAsByte("./config.json"))
  if configJSONErr != nil {
    panic(configJSONErr)
  }

  // Get our Music Directory from out config JSON
  musicDir, doesExist := configJSON.Path("musicFilesPath").Data().(string)
  if doesExist == false {
    panic("config.json: musicFilesPath Does Not Exist in JSON")
  }

  // Get our playlist JSON, and set an array on songs
  playlistJSON, playlistJSONErr := gabs.ParseJSON(renderer.ReadFileAsByte("./templates/pages/playlist/playlist.json"))
  if playlistJSONErr != nil {
    panic(playlistJSONErr)
  }
  playlistJSON.ArrayP("songs");

  // Iterate through the directory (Already in Lexical order, may need to sort the song JSON by Artist, and then song)
  filepath.Walk(musicDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
      // Open the file for reading
      file, err := os.Open(path)
      if err != nil {
      	panic(err)
      }

      // Send the file to the meta data tagger
      metaData, err := tag.ReadFrom(file)
      if err != nil {
      	panic(err)
      }

      songTitle := metaData.Title()
      songArtist := metaData.Artist()

      if len(songTitle) > 0 && len(songTitle) > 0 {
        // Create the Song JSON
        songJSON := gabs.New()
        songJSON.SetP(songArtist, "artist")
        songJSON.SetP(songTitle, "title")

        // Add the song JSON to the playlist songs array
        playlistJSON.ArrayAppend(songJSON.Data(), "songs")
      }
		}
		return nil
	})

c.Data(200, "text/html; charset=utf-8",
  []byte(renderer.RenderFromRawJSON("./templates/pages/playlist/playlist.html", []byte(playlistJSON.String()))))
}
