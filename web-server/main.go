package main

import (
	"github.com/calebmchenry/mindful/web-server/database"
	"github.com/calebmchenry/mindful/web-server/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.Init()
	s := server.New()
	s.Serve()
}
