package main

import (
	"github.com/DanielCorreiaPina/realstateAPI/app"
	"github.com/DanielCorreiaPina/realstateAPI/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
