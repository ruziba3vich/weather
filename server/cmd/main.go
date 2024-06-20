package main

import (
	"log"

	"github.com/ruziba3vich/local_weather_server/api"
	"github.com/ruziba3vich/local_weather_server/internal/service"
	"github.com/ruziba3vich/local_weather_server/internal/storage"
)

func main() {
	api := api.New(service.New(storage.New()))
	log.Fatal(api.RUN("7777"))
}
