package client

import (
	"context"
	"log"

	"github.com/johnldev/go-expert/chanllenge-1/internal/client/services"
)

func GetDollarPrice() {
	ctx, cancel := context.WithCancel(context.Background())

	result, err := services.GetDolarPrice(ctx)
	if err != nil {
		log.Fatalln("Error to get dollar price: ", err.Error())
	}

	services.SaveDollarPrice(*result)
	cancel()

}
