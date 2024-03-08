package services

import (
	"fmt"
	"net/http"
)

type FipeService struct{}

func (s *FipeService) GetBrands() *http.Response {
	url := "https://parallelum.com.br/fipe/api/v1/carros/marcas"

	res, _ := http.Get(url)

	return res
}

func (s *FipeService) GetModels(brand string) *http.Response {
	url := fmt.Sprintf("https://parallelum.com.br/fipe/api/v1/carros/marcas/%s/modelos", brand)

	res, _ := http.Get(url)

	return res
}

func (s *FipeService) GetYears(brand string, model string) *http.Response {
	url := fmt.Sprintf("https://parallelum.com.br/fipe/api/v1/carros/marcas/%s/modelos/%s/anos", brand, model)

	res, _ := http.Get(url)

	return res
}

func (s *FipeService) GetFull(brand string, model string, year string) *http.Response {
	url := fmt.Sprintf("https://parallelum.com.br/fipe/api/v1/carros/marcas/%s/modelos/%s/anos/%s", brand, model, year)

	res, _ := http.Get(url)

	return res
}
