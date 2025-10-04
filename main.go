package main

import (
	"github.com/HanmaDevin/image-processing/database"
	"github.com/HanmaDevin/image-processing/server"
)

func main() {
	s := server.NewServer()
	database.InitDB()
	server.StartServer(s)
}
