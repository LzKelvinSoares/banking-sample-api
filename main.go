package main

import (
	"banking/app"

	"github.com/LzKelvinSoares/lzk-go-lib/logger"
)

func main() {
	// log.Println("Starting App...")
	logger.Info("Starting App")
	app.Start()
}
