package api

import (
	"log"
	"net"

	genprotos "github.com/ruziba3vich/local_weather/genprotos/protos"
	"google.golang.org/grpc"
)

type API struct {
	service genprotos.WeatherServiceServer
}

func New(service genprotos.WeatherServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	genprotos.RegisterWeatherServiceServer(serverRegisterer, a.service)
	log.Println("server has been started on", port)

	return serverRegisterer.Serve(listener)
}
