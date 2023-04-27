package main

import "notetaking-go/internals/app"

func main() {
	app := app.New("api/v1", ".env", "note.db")
	app.RunServer()
}
