package main

import (
	"github.com/johnldev/go-expert/chanllenge-1/internal/server"
	"github.com/johnldev/go-expert/chanllenge-1/internal/server/database"
)

func main() {
	server.StartHttpServer(database.NewConnection())
}
