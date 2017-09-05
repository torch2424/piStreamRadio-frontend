package main

import "fmt"
import "encoding/json"

//3P libs
import "github.com/gin-gonic/gin"
import "github.com/Jeffail/gabs"
import "github.com/aymerick/raymond"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// Some sample json, TODO: need to read froma  file
		sampleJson := []byte(`{
			"title": "test",
			"body": "ayyeee lmao",
			"deep": {
				"rolling": "in the deep"
			}
		}`)

		// Digest JSON in gabs to allow editing
		jsonParsed, _ := gabs.ParseJSON(sampleJson)
		jsonParsed.SetP("I changed the body! lmaooo ayyeee", "body");
		// Unmarshal the json
		// https://gobyexample.com/json
		var jsonMarshal map[string]interface{}
		err := json.Unmarshal([]byte(jsonParsed.String()), &jsonMarshal);
		if err != nil {
			panic(err)
		}

		// Our example template need to TODO: read from a file
		tpl := `<div class="entry">
		hello testing!
		 <h1>{{title}}</h1>
		 <div class="body">
			 {{body}}
			 <div>
			 	{{deep}}
			 </div>
			 <div>
			 	{{deep.rolling}}
			 </div>
		 </div>
		</div>
		`

		// Render our handlebars Template
		result, err := raymond.Render(tpl, jsonMarshal)
    if err != nil {
        panic("Please report a bug :)")
    }

		fmt.Println("Rendered /")

		// Return our html
		c.Data(200, "text/html; charset=utf-8", []byte(result))
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
