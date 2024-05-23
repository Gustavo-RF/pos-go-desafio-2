package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Gustavo-RF/desafio-go-2/brasilapi"
	"github.com/Gustavo-RF/desafio-go-2/viacep"
	"github.com/paemuri/brdoc"
)

type CepResponse struct {
	Origin        string
	Timespan      time.Duration
	Cep           string
	Address       string
	Neighbourhood string
	City          string
	State         string
}

func main() {

	cepArgument := os.Args[1:]

	if len(cepArgument) < 1 {
		fmt.Println("Send a zip code via argument.\nExample: go run main.go mycep")
		return
	}

	if len(cepArgument) > 1 {
		fmt.Println("Send only a zip code via argument.\nExample: go run main.go mycep")
		return
	}

	cep := os.Args[1]

	if !brdoc.IsCEP(cep) {
		fmt.Println("Invalid zip code")
		return
	}

	c1 := make(chan CepResponse)
	c2 := make(chan CepResponse)

	go func() {
		start := time.Now()
		brasilApi, err := brasilapi.Fetch(cep)
		if err != nil {
			panic(err)
		}
		end := time.Now()

		diff := end.Sub(start)

		cepResponseBrailApi := CepResponse{
			Origin:        "Brasil api",
			Timespan:      diff,
			Cep:           brasilApi.Cep,
			Address:       brasilApi.Street,
			Neighbourhood: brasilApi.Neighborhood,
			City:          brasilApi.City,
			State:         brasilApi.State,
		}

		c1 <- cepResponseBrailApi
	}()

	go func() {
		start := time.Now()
		viacep, err := viacep.Fetch(cep)
		if err != nil {
			panic(err)
		}

		end := time.Now()

		diff := end.Sub(start)

		cepResponseViaCep := CepResponse{
			Origin:        "Viacep",
			Timespan:      diff,
			Cep:           viacep.Cep,
			Address:       viacep.Logradouro,
			Neighbourhood: viacep.Bairro,
			City:          viacep.Localidade,
			State:         viacep.Uf,
		}

		c2 <- cepResponseViaCep
	}()

	select {
	case msg1 := <-c1:
		fmt.Printf("Received from %s in %s.\nAddress: %s - %s - %s / %s\n", msg1.Origin, msg1.Timespan, msg1.Address, msg1.Neighbourhood, msg1.City, msg1.State)
	case msg2 := <-c2:
		fmt.Printf("Received from %s in %s.\nAddress: %s - %s - %s / %s\n", msg2.Origin, msg2.Timespan, msg2.Address, msg2.Neighbourhood, msg2.City, msg2.State)
	case <-time.After(1 * time.Second):
		fmt.Println("Not received. Timeout")
	}
}
