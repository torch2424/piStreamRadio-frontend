# piStreamRadio-frontend

![piStreamRadio Frontend Screenshot](https://files.aaronthedev.com/$/5xdkp)

A frontend web client for piStreamRadio. Using [Gin](https://github.com/gin-gonic/gin), [AMP](https://github.com/ampproject/amphtml), and [MustacheJS with raymond](https://github.com/aymerick/raymond). Styling is done with [AMP Start](https://www.ampstart.com/) and [Animate.css](https://daneden.github.io/animate.css/). A live example can be found at [Galaxy Noise Radio](https://galaxynoiseradio.com/).

**Features:**
* Embed of live stream on home page
* Fully configurable using the `.json` files
* Playlist of songs on the /playlist route

# Getting Started (Installation)

1. [Install Go](https://golang.org/doc/install) This project was built on Go 1.9.

2. Clone the project:

```
git clone https://github.com/torch2424/piStreamRadio-frontend.git
cd piStreamRadio-frontend
```

3. Edit the `config.json` and put your information in the json values:

```
vim config.json
```

4. Run the Project with:

```
go run main.go
```

# Customizing The Site

Please feel free to edit any of the files to your liking. For more simple users, you may mostly want to play with the `.json` files under the `templates/pages` directory. More advanced users may want to edit both the `.json` and `.html` files under the `templates/pages` directory. New Routes can be added by creating directories under the `templates/pages` directory, and registering the route in the `main.go` file.

# Contributing

Feel free to fork the project, open up a PR, and give any contributions! I'd suggest opening an issue first however, just so everyone is aware and can discuss the proposed changes.

# LICENSE

LICENSE under [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
