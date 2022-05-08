package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DanielCorreiaPina/realstateAPI/service"
	"github.com/gorilla/mux"
)

type HouseHandlers struct {
	service service.HouseService
}

func (hh *HouseHandlers) getAllHouses(w http.ResponseWriter, r *http.Request) {
	houses, _ := hh.service.GetAllHouse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

func (hh *HouseHandlers) getHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["house_id"]

	house, err := hh.service.GetHouse(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(house)
	}
}
