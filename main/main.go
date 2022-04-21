package main

import (
	"log"
	"net/http"

	"github.com/Diegoplas/technicalChallenge2022/configuration"
	"github.com/Diegoplas/technicalChallenge2022/route"
	"github.com/gorilla/handlers"
)

func main() {
	router := route.GetRouter()
	methods := handlers.AllowedMethods([]string{http.MethodGet})
	log.Fatal(http.ListenAndServe(configuration.Port, handlers.CORS(methods)(router)))
}
