package server

import (
	"fmt"
	"net/http"

	"github.com/johnldev/go-expert/chanllenge-1/internal/server/handlers"
)

const (
	PORT = 8080
)

func StartHttpServer() {

	server := http.NewServeMux()

	server.HandleFunc("/cotacao", handlers.DolarPriceHandler)

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), server)

}
