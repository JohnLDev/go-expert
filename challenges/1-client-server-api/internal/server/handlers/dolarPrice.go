package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/johnldev/go-expert/chanllenge-1/internal/server/repositories"
	"github.com/johnldev/go-expert/chanllenge-1/internal/server/services"
)

func DolarPriceHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("DolarPriceHandler")
		// external api
		result, err := services.GetDollarPrice(r.Context())
		if err != nil {
			fmt.Println("Error to get dollar price: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// save in database

		repo := repositories.NewDolarPriceRepository(db, r.Context())
		err = repo.Save(result.Usdbrl)
		if err != nil {
			fmt.Println("Error to save dollar price: ", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// response
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(result.Usdbrl)
		if err != nil {
			panic(err)
		}
	}
}
