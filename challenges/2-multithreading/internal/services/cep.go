package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/johnldev/go-expert/chanllenge-2/internal/utils"
)

type cepService struct {
	ctx context.Context
}

func (s cepService) GetCep(cep string) {

	cdnUrl := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
	viaCepUrl := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", strings.Replace(cep, "-", "", 1))

	resultViaCep := make(chan []byte)
	resultCdn := make(chan []byte)
	defer close(resultViaCep)
	defer close(resultCdn)

	ctx, cancel := context.WithTimeout(s.ctx, time.Second*1)
	defer cancel()

	go func() {
		response, _ := utils.RequestWithContext(ctx, cdnUrl)
		resultCdn <- response
	}()

	go func() {
		response, _ := utils.RequestWithContext(ctx, viaCepUrl)
		resultViaCep <- response
	}()

	select {
	case result := <-resultCdn:
		fmt.Println("Response from cdn")
		fmt.Println(string(result))
	case result := <-resultViaCep:
		fmt.Println("Response from viacep")
		fmt.Println(string(result))
	case <-ctx.Done():
		fmt.Println("Timeout on request")
		fmt.Println(ctx.Err())
	}
}

func NewCepService(ctx context.Context) cepService {
	return cepService{ctx: ctx}
}
