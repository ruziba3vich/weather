package storage

import (
	"fmt"

	genprotos "github.com/ruziba3vich/local_weather_server/genprotos/protos"
)

type (
	Storage struct {
		db *DB
	}
)

func New() *Storage {
	return &Storage{
		db: GetDB(),
	}
}

func (s *Storage) GetWeatherByCountryName(req *genprotos.GetWeatherByCountryNameReq) float32 {
	return s.db.GetRandomTemperature()
}

func (s *Storage) FindCountry(req *genprotos.GetWeatherByCountryNameReq) error {
	if !s.db.FindCountry(req.CountryName) {
		return fmt.Errorf("country %s has not been found", req.CountryName)
	}
	return nil
}
