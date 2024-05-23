package brasilapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type BrasilApiDto struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Type         string `json:"type"`
}

func Fetch(cep string) (*BrasilApiDto, error) {
	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/cep/v1/"+cep, nil)

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

	var data BrasilApiDto
	err = json.Unmarshal(res, &data)

	if err != nil {
		return nil, err
	}

	if data.Type != "" {
		return nil, http.ErrNotSupported
	}

	return &data, nil
}
