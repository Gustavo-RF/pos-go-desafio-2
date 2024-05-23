package viacep

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViacepDto struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func Fetch(cep string) (*ViacepDto, error) {
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data ViacepDto
	err = json.Unmarshal(res, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
