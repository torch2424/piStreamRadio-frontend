package renderer

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

// 3P libs
import (
  "github.com/aymerick/raymond"
  "github.com/Jeffail/gabs"
)

// Read file into string
// https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable-in-golang

func appendSiteJSON(rawJSON []byte) []byte {

  // Parse our passed JSON
  rawJSONParsed, rawParseErr := gabs.ParseJSON(rawJSON)
  if rawParseErr != nil {
    panic(rawParseErr)
  }

  // Parse our config JSON
  configJSONParsed, configParseErr := gabs.ParseJSON(ReadFileAsByte("./config.json"))
  if configParseErr != nil {
    panic(configParseErr)
  }

  // Set the config JSON to the config param in the json itself
  rawJSONParsed.SetP(configJSONParsed.Data(), "config");

  // Return the new raw JSON
  return []byte(rawJSONParsed.String())
}

// ReadFileAsByte Helper Function to read the file at the path and return as []byte
func ReadFileAsByte(filePath string) []byte {
  file, readErr := ioutil.ReadFile(filePath);
  if readErr != nil {
    panic(readErr)
  }

  return file
}

// RenderFromPath funnction to render the specified template path with the passed json path
func RenderFromPath(templatePath string, jsonPath string) string {

  // Read the template as a byte[]
  template := ReadFileAsByte(templatePath)
  // Read the Json as a byte[]
  rawJSON := ReadFileAsByte(jsonPath)

  return Render(template, rawJSON);
}

// RenderFromRawJSON funnction to render the specified template path with the passed raw json byte array
func RenderFromRawJSON(templatePath string, rawJSON []byte) string {
  // Read the template as a byte[]
  template := ReadFileAsByte(templatePath)

  return Render(template, rawJSON);
}

// RenderFromRawTemplate funnction to render the specified raw template with the passed json path
func RenderFromRawTemplate(template []byte, jsonPath string) string {
  // Read the Json as a byte[]
  rawJSON := ReadFileAsByte(jsonPath)

  return Render(template, rawJSON);
}

// Render function to render the specified raw template with the passed raw json
func Render(template []byte, rawJSON []byte) string {

  fmt.Println("Rendering Template...");

  // Add the site config json to the passed json for partials
  rawJSON = appendSiteJSON(rawJSON)

  // Unmarshal the json
  // https://gobyexample.com/json
  var unMarshalJSON map[string]interface{}
  unMarshalErr := json.Unmarshal(rawJSON, &unMarshalJSON)
  if unMarshalErr != nil {
    panic(unMarshalErr)
  }

  // Render our handlebars Template
  renderedTemplate, renderErr := raymond.Render(string(template), unMarshalJSON);
  if renderErr != nil {
      panic(renderErr)
  }

  return renderedTemplate
}
