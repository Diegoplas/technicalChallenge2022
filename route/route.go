package route

import (
	"net/http"

	"github.com/Diegoplas/technicalChallenge2022/controller"
	"github.com/gorilla/mux"
)

func GetRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/hello", controller.HelloWize).Methods(http.MethodGet)
	router.HandleFunc("/play-cards", controller.PlayCards).Methods(http.MethodGet)
	return
}
