package app

import (
	"encoding/json"
	"net/http"

	"github.com/DanielCorreiaPina/realstateAPI/service"
	"github.com/gorilla/mux"
)

type HouseHandlers struct {
	service service.HouseService
}

func (hh *HouseHandlers) getAllHouses(w http.ResponseWriter, r *http.Request) {
	houses, err := hh.service.GetAllHouse()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, houses)
	}
}

func (hh *HouseHandlers) getHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["house_id"]

	house, err := hh.service.GetHouse(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, house)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
