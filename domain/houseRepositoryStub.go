package domain

type HouseRepositoryStub struct {
	houses []House
}

func (s HouseRepositoryStub) FindAll() ([]House, error) {
	return s.houses, nil
}

func NewHouseRepositoryStub() HouseRepositoryStub {
	houses := []House{
		{1, "apartament", 3, 3, 90, 2022, "new", "a", true, false, false, true, true, "Address 1"},
		{2, "apartament", 2, 1, 85, 2015, "second hand", "b", true, false, false, false, true, "Address 2"},
		{3, "apartament", 3, 2, 90, 2021, "second hand", "a", true, false, false, true, false, "Address 3"},
	}
	return HouseRepositoryStub{houses}
}
