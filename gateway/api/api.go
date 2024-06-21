package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/local_weather_gateway/api/handler"
	"google.golang.org/grpc"
)

type API struct {
	logger *log.Logger
}

func New(logger *log.Logger) *API {
	return &API{
		logger: logger,
	}
}

func (a *API) RUN(port, host string, connection *grpc.ClientConn) error {
	handler := handler.New(host, connection, a.logger)

	router := gin.Default()
	router.GET("/get/by/country/name", handler.GetWeatherByCountryName)

	return router.Run(":" + port)
}
