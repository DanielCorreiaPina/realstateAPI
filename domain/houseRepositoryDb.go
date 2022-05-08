package domain

import (
	"database/sql"
	"fmt"
	"log"
	os "os"
	"strconv"
	"time"

	"github.com/DanielCorreiaPina/realstateAPI/errs"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type HouseRepositoryDb struct {
	client *sql.DB
}

func (d HouseRepositoryDb) FindAll() ([]House, error) {
	findAllSql := "select * from house"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying house table " + err.Error())
		return nil, err
	}

	houses := make([]House, 0)
	for rows.Next() {
		var h House
		err := rows.Scan(&h.Id, &h.ConstructionType, &h.NumberBedrooms, &h.NumberBathrooms, &h.Area, &h.ConstructionYear, &h.Condition, &h.EnergeticCertificate, &h.GarageSpots, &h.ParkingLots, &h.SwimmingPools, &h.Elevators, &h.FullyEquipped, &h.Address)
		if err != nil {
			log.Println("Error while scanning houses " + err.Error())
			return nil, err
		}
		houses = append(houses, h)
	}
	return houses, nil
}

func (d HouseRepositoryDb) FindById(id string) (*House, *errs.AppError) {
	houseSql := "select * from house where id = $1"

	row := d.client.QueryRow(houseSql, id)
	var h House
	err := row.Scan(&h.Id, &h.ConstructionType, &h.NumberBedrooms, &h.NumberBathrooms, &h.Area, &h.ConstructionYear, &h.Condition, &h.EnergeticCertificate, &h.GarageSpots, &h.ParkingLots, &h.SwimmingPools, &h.Elevators, &h.FullyEquipped, &h.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("House not found")
		} else {
			log.Println("Error while scanning house " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &h, nil
}

func NewHouseRepositoryDb() HouseRepositoryDb {
	host := os.Getenv("POSTGRES_HOST")
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Println("Error while reading POSTGRES_PORT environment variable " + err.Error())
		panic(err)
	}
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return HouseRepositoryDb{client}
}
