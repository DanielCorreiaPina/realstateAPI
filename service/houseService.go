package service

import (
	"github.com/DanielCorreiaPina/realstateAPI/domain"
	"github.com/DanielCorreiaPina/realstateAPI/errs"
)

type HouseService interface {
	GetAllHouse() ([]domain.House, error)
	GetHouse(string) (*domain.House, *errs.AppError)
}

type DefaultHouseService struct {
	repo domain.HouseRepository
}

func (s DefaultHouseService) GetAllHouse() ([]domain.House, error) {
	return s.repo.FindAll()
}

func (s DefaultHouseService) GetHouse(id string) (*domain.House, *errs.AppError) {
	return s.repo.FindById(id)
}

func NewHouseService(repository domain.HouseRepository) DefaultHouseService {
	return DefaultHouseService{repository}
}
