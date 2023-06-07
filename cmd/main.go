package main

import (
	"github.com/DmitryOdintsov/workingWithGit/internal/server"
	"github.com/DmitryOdintsov/workingWithGit/internal/server/handlers"
)

func main() {
	hand := handlers.NewHandlers()
	server.Run(hand)
}
