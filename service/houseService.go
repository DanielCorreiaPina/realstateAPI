package service

import "github.com/DanielCorreiaPina/realstateAPI/domain"

type HouseService interface {
	GetAllHouse() ([]domain.House, error)
}

type DefaultHouseService struct {
	repo domain.HouseRepository
}

func (s DefaultHouseService) GetAllHouse() ([]domain.House, error) {
	return s.repo.FindAll()
}

func NewHouseService(repository domain.HouseRepository) DefaultHouseService {
	return DefaultHouseService{repository}
}
