package main

import app "github.com/eugene-krivtsov/idler/internal/app/facade-app"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
