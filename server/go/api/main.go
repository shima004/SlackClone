package main

import (
	"Slack/server"
)

func main() {
	e := server.NewServer()
	e.Logger.Fatal(e.Start(":8080"))
}
