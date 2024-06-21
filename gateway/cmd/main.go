package main

import (
	"log"
	"os"

	"github.com/ruziba3vich/local_weather_gateway/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:7777", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	api := api.New(log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile))

	log.Fatal(api.RUN("8888", "localhost:7777", conn))
}
