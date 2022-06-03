package main

import app "github.com/eugene-krivtsov/idler/internal/app/chat-service"

const (
	configPath = "config/config"
)

func main() {
	app.Run(configPath)
}
