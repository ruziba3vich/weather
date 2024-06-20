package main

import (
	"log"

	"github.com/ruziba3vich/local_weather/api"
	"github.com/ruziba3vich/local_weather/internal/service"
	"github.com/ruziba3vich/local_weather/internal/storage"
)

func main() {
	api := api.New(service.New(storage.New()))
	log.Fatal(api.RUN("7777"))
}
