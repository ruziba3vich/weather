package service

import (
	"log"
	"os"
	"time"

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
		s.logger.Println(err)
		return err
	}

	for range 10 {
		response := &genprotos.GetWeatherByCountryNameRes{
			CountryName: req.CountryName,
			Temperature: s.storgae.GetWeatherByCountryName(req),
		}
		// log.Println("------------------------------------------got out------------------------------------------")
		streamer.Send(response)
		log.Println(response)
		time.Sleep(time.Second)
	}
	return nil
}
