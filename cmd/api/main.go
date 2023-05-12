package main

import "notetaking-go/internals/app"

func main() {
	app := app.New("api/v1", ".env")
	app.RunServer()
}
