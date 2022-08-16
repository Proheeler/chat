package main

import (
	"chat/cmd"
	"chat/internal/logic"
)

func main() {
	app := cmd.App{
		H: logic.NewHub(),
	}
	app.Run()
}
