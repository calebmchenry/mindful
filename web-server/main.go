package main

import "github.com/calebmchenry/mindful/web-server/server"

func main() {
	s := server.New()
	s.Serve()
}
