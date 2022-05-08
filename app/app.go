package app

import (
	"log"
	"net/http"

	"github.com/DanielCorreiaPina/realstateAPI/domain"
	"github.com/DanielCorreiaPina/realstateAPI/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	//hh := HouseHandlers{service.NewHouseService(domain.NewHouseRepositoryStub())}
	hh := HouseHandlers{service.NewHouseService(domain.NewHouseRepositoryDb())}

	//define routes
	router.HandleFunc("/houses", hh.getAllHouses).Methods(http.MethodGet)
	router.HandleFunc("/houses/{house_id:[0-9]+}", hh.getHouse).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
