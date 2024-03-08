package controllers

import (
	"fipe/services"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type FipeController struct {
	fipe services.FipeService
}

func NewFipeController() *FipeController {
	controller := FipeController{
		fipe: services.FipeService{},
	}

	return &controller
}

func (c *FipeController) getRequestBody(r *http.Response) string {
	bodyBytes, _ := io.ReadAll(r.Body)

	return string(bodyBytes)
}

func (c *FipeController) Brands(w http.ResponseWriter, r *http.Request) {
	res := c.fipe.GetBrands()

	fmt.Fprint(w, c.getRequestBody(res))
}

func (c *FipeController) Models(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	brand := vars["brand"]

	res := c.fipe.GetModels(brand)

	fmt.Fprint(w, c.getRequestBody(res))
}

func (c *FipeController) Years(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	brand := vars["brand"]
	model := vars["model"]

	res := c.fipe.GetYears(brand, model)

	fmt.Fprint(w, c.getRequestBody(res))
}

func (c *FipeController) Full(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	brand := vars["brand"]
	model := vars["model"]
	year := vars["year"]

	res := c.fipe.GetFull(brand, model, year)

	fmt.Fprint(w, c.getRequestBody(res))
}
