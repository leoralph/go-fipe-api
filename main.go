package main

import (
	"fipe/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	router := mux.NewRouter()

	fipeController := controllers.NewFipeController()

	router.HandleFunc("/marcas", fipeController.Brands).Methods("GET")
	router.HandleFunc("/modelos/{brand}", fipeController.Models).Methods("GET")
	router.HandleFunc("/anos/{brand}/{model}", fipeController.Years).Methods("GET")
	router.HandleFunc("/valor/{brand}/{model}/{year}", fipeController.Full).Methods("GET")

	return router
}

func main() {
	http.ListenAndServe(":8080", getRouter())
}
