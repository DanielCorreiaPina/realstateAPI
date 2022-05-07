package app

import (
	"encoding/json"
	"net/http"

	"github.com/DanielCorreiaPina/realstateAPI/service"
)

type HouseHandlers struct {
	service service.HouseService
}

func (hh *HouseHandlers) getAllHouses(w http.ResponseWriter, r *http.Request) {
	houses, _ := hh.service.GetAllHouse()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}
