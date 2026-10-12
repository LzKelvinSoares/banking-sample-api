package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	// log.Println("Starting App...")
	logger.Info("Starting App")
	app.Start()
}
