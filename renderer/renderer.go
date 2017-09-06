package renderer

import (
  "io/ioutil"
  "encoding/json"
)

// 3P libs
import (
  "github.com/aymerick/raymond"
)

// Read file into string
// https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable-in-golang

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
