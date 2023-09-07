package main

import (
	"os"
)

func main() {
	err := ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable is not set")
	}
	app := NewBotApp(token)

	app.Run()
}
