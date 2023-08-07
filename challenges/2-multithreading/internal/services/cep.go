package services

import (
	"context"
	"fmt"

	"github.com/johnldev/go-expert/chanllenge-2/internal/utils"
)

type cepService struct {
	ctx context.Context
}

func (s cepService) GetCep(cep string) {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	cdnUrl := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	viaCepUrl := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	result := make(chan []byte)
	defer close(result)

	go func() {
		response, err := utils.RequestWithContext(ctx, cdnUrl)
		if err != nil {
			fmt.Printf("Error on request to cdn: %s\n", err.Error())
			return
		}
		result <- response
		fmt.Println("Response from cdn")
		cancel()
	}()

	go func() {
		response, err := utils.RequestWithContext(ctx, viaCepUrl)
		if err != nil {
			fmt.Printf("Error on request to viacep: %s\n", err.Error())
			return
		}
		result <- response
		fmt.Println("Response from viacep")
		cancel()
	}()

	fmt.Println(string(<-result))
}

func NewCepService(ctx context.Context) cepService {
	return cepService{ctx: ctx}
}
