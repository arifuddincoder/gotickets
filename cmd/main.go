package main

import (
	"gotickets/internal/config"
	"gotickets/internal/server"
)

func main() {
	cfg := config.LoadEnv()

	db := config.InitDB(cfg)

	server.Start(db, cfg)
}
