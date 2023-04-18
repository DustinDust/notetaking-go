package main

import "notetaking-go/pkg/app"

func main() {
	app := app.New("api/v1")
	app.RunServer()
}
