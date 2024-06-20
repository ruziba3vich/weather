package service

import (
	"log"
	"os"

	genprotos "github.com/ruziba3vich/local_weather_server/genprotos/protos"
	"github.com/ruziba3vich/local_weather_server/internal/storage"
)

type (
	service struct {
		logger  *log.Logger
		storgae *storage.Storage
		genprotos.UnimplementedWeatherServiceServer
	}
)

func New(storgae *storage.Storage) *service {
	return &service{
		storgae: storgae,
		logger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (s *service) GetWeatherByCountryName(req *genprotos.GetWeatherByCountryNameReq, streamer genprotos.WeatherService_GetWeatherByCountryNameServer) error {
	s.logger.Println("SERVER RECIEVED A REQUEST WITH", req.CountryName)
	if err := s.storgae.FindCountry(req); err != nil {
		return err
	}

	for range 50 {
		ch := make(chan float32)
		s.storgae.GetWeatherByCountryName(req, ch)
		streamer.Send(
			&genprotos.GetWeatherByCountryNameRes{
				CountryName: req.CountryName,
				Temperature: <-ch,
			})
	}
	return nil
}
