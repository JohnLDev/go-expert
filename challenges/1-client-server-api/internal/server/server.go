package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/johnldev/go-expert/chanllenge-1/internal/server/handlers"
)

const (
	PORT = 8080
)

func StartHttpServer(db *sql.DB) {

	server := http.NewServeMux()

	server.HandleFunc("/cotacao", handlers.DolarPriceHandler(db))

	http.ListenAndServe(fmt.Sprintf(":%d", PORT), server)

}
