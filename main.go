package main

import (
	"./app"
)

func main() {
	app := &app.App{}
	app.Run(":8080")
}
